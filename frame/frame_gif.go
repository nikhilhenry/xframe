package frame

import (
	"bytes"
	"golang.org/x/image/draw"
	"image"
	"image/color"
	pallete "image/color/palette"
	"image/gif"
	"io"
)

func GenerateFrameWithBezelGIF(w io.Writer, imageGif gif.GIF) error {

	//decode images
	imageFrames := imageGif.Image

	// temporary array with palette images [store the processed buffers]
	framedImages := make([]*image.Paletted, 5)
	// @todo encodes for all image buffer not just 5
	for index, imageFrame := range imageFrames[0:5] {
		imageBuf := bytes.Buffer{}

		const imageWidth = 1170
		const imageHeight = 2532
		scaledDstImage := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
		draw.NearestNeighbor.Scale(scaledDstImage, scaledDstImage.Bounds(), imageFrame, imageFrame.Bounds(), draw.Over, nil)
		err := GenerateFrameWithBezel(&imageBuf, scaledDstImage)
		if err != nil {
			return err
		}

		// convert to palette image for processing
		framedImage, _, err := image.Decode(&imageBuf)
		if err != nil {
			return err
		}
		framedImages[index] = imageToPaleted(framedImage, pallete.Plan9)
	}

	//encode the buffer as a  gif
	processedGIF := gif.GIF{Image: framedImages, Delay: imageGif.Delay[0:5]}
	err := gif.EncodeAll(w, &processedGIF)
	if err != nil {
		return err
	}

	return nil
}

func imageToPaleted(img image.Image, pallete color.Palette) *image.Paletted {
	bounds := img.Bounds()
	palettedImage := image.NewPaletted(bounds, pallete)
	draw.FloydSteinberg.Draw(palettedImage, bounds, img, image.Point{})
	return palettedImage
}
