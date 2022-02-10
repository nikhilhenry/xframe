package frame

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	pallete "image/color/palette"
	"image/draw"
	"image/gif"
	"io"
)

func GenerateFrameWithBezelGIF(w io.Writer, imageGif gif.GIF) error {

	//decode images
	imageFrames := imageGif.Image

	framedImages := make([]*image.Paletted, 5)

	for index, imageFrame := range imageFrames[0:5] {
		imageBuf := bytes.Buffer{}
		fmt.Println("going to generate image")

		err := GenerateFrameWithBezel(&imageBuf, imageFrame)
		if err != nil {
			return err
		}
		fmt.Println("image framed")
		framedImage, _, err := image.Decode(&imageBuf)
		if err != nil {
			return err
		}
		fmt.Println("image encoded")
		framedImages[index] = imageToPaleted(framedImage, pallete.Plan9)
	}

	//create gif
	fmt.Println(len(framedImages))
	fmt.Println(len(imageGif.Delay[0:5]))
	processedGIF := gif.GIF{Image: framedImages, Delay: imageGif.Delay[0:5]}
	//encode all the images
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
