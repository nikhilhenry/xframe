package utils

import (
	"image"
	"image/color/palette"
	"image/gif"
	"image/png"
	"io"
)

func EncodePNG(w io.Writer) func(image.Image) error {
	return func(drawableImg image.Image) error {
		if err := png.Encode(w, drawableImg); err != nil {
			return err
		}
		return nil
	}
}

type result struct {
	index int
	image *image.Paletted
}

func EncodeGIF(w io.Writer, delay []int) func([]image.Image) error {
	return func(imgs []image.Image) error {
		palettedImgs := make([]*image.Paletted, len(imgs))
		resultsChannel := make(chan result)
		// convert image.Image to image.Paletted concurrently
		for i, img := range imgs {
			go func(i int, img image.Image) {
				resultsChannel <- result{index: i, image: ImageToPaletted(img, palette.Plan9)}
			}(i, img)
		}
		for i := 0; i < len(imgs); i++ {
			r := <-resultsChannel
			palettedImgs[r.index] = r.image
		}
		processedGIF := gif.GIF{Image: palettedImgs, Delay: delay}
		if err := gif.EncodeAll(w, &processedGIF); err != nil {
			return err
		}
		return nil
	}
}
