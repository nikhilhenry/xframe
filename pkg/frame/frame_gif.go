package frame

import (
	"bytes"
	bezel "github.com/nikhilhenry/xframe/internal/bezel"
	"golang.org/x/image/draw"
	"image"
	"image/color"
	pallete "image/color/palette"
	"image/gif"
	"io"
)

type result struct {
	index int
	image *image.Paletted
}

func GenerateFrameWithBezelGIF(w io.Writer, imageGif gif.GIF) error {

	//decode images
	imageFrames := imageGif.Image

	// temporary array with palette images [store the processed buffers]
	framedImages := make([]*image.Paletted, len(imageFrames))
	// channel to support concurrency
	resultsChannel := make(chan result)

	const imageWidth = 1170
	const imageHeight = 2532

	deviceBezel := bezel.New(bezel.Iphone13Pro)

	for index, img := range imageFrames {
		go func(i int, imageFrame *image.Paletted) error {

			imageBuf := bytes.Buffer{}
			scaledDstImage := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
			draw.NearestNeighbor.Scale(scaledDstImage, scaledDstImage.Bounds(), imageFrame, imageFrame.Bounds(), draw.Over, nil)
			err := GenerateFrameWithBezel(&imageBuf, *deviceBezel, scaledDstImage)
			if err != nil {
				return err
			}

			// convert to palette image for processing
			framedImage, _, err := image.Decode(&imageBuf)
			if err != nil {
				return err
			}
			resultsChannel <- result{index: i, image: imageToPaletted(framedImage, pallete.Plan9)}

			return nil
		}(index, img)
	}

	for i := 0; i < len(imageFrames); i++ {
		r := <-resultsChannel
		framedImages[r.index] = r.image
	}

	//encode the buffer as a  gif
	processedGIF := gif.GIF{Image: framedImages, Delay: imageGif.Delay}
	err := gif.EncodeAll(w, &processedGIF)
	if err != nil {
		return err
	}

	return nil
}

func imageToPaletted(img image.Image, palette color.Palette) *image.Paletted {
	bounds := img.Bounds()
	palettedImage := image.NewPaletted(bounds, palette)
	draw.FloydSteinberg.Draw(palettedImage, bounds, img, image.Point{})
	return palettedImage
}
