package frame

import (
	"golang.org/x/image/draw"
	"image"
)

type result struct {
	index int
	image image.Image
}

func GenerateSequence(encode func([]image.Image) error, overlay overlay, screenImages []*image.Image) error {
	// temporary array with palette images [store the processed buffers]
	framedImages := make([]image.Image, len(screenImages))
	// channel to support concurrency
	resultsChannel := make(chan result)

	const imageWidth = 1170
	const imageHeight = 2532

	for index, img := range screenImages {
		go func(i int, imageFrame *image.Image) error {

			scaledDstImage := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
			draw.NearestNeighbor.Scale(scaledDstImage, scaledDstImage.Bounds(), *imageFrame, (*imageFrame).Bounds(), draw.Over, nil)
			err := Generate(func(rgba *image.RGBA) error {
				resultsChannel <- result{index: i, image: rgba}
				return nil
			}, overlay, scaledDstImage)
			if err != nil {
				return err
			}
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
