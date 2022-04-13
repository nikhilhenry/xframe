package frame

import (
	"bytes"
	"github.com/nikhilhenry/xframe/internal/utils"
	"golang.org/x/image/draw"
	"image"
	"image/color/palette"
)

// @todo accept array of image.Image instead of image.Paletted

func GenerateSequence(encode func([]*image.Paletted) error, overlay overlay, screenImages []*image.Paletted) error {
	// temporary array with palette images [store the processed buffers]
	framedImages := make([]*image.Paletted, len(screenImages))
	// channel to support concurrency
	resultsChannel := make(chan result)

	const imageWidth = 1170
	const imageHeight = 2532

	for index, img := range screenImages {
		go func(i int, imageFrame *image.Paletted) error {

			imageBuf := bytes.Buffer{}
			scaledDstImage := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
			draw.NearestNeighbor.Scale(scaledDstImage, scaledDstImage.Bounds(), imageFrame, imageFrame.Bounds(), draw.Over, nil)
			err := Generate(utils.EncodePNG(&imageBuf), overlay, scaledDstImage)
			if err != nil {
				return err
			}
			//@todo palette image without png encoding for fast performance
			// convert to palette image for processing
			framedImage, _, err := image.Decode(&imageBuf)
			if err != nil {
				return err
			}
			resultsChannel <- result{index: i, image: imageToPaletted(framedImage, palette.Plan9)}

			return nil
		}(index, img)
	}

	for i := 0; i < len(screenImages); i++ {
		r := <-resultsChannel
		framedImages[r.index] = r.image
	}
	// encode images
	if err := encode(framedImages); err != nil {
		return err
	}
	return nil
}
