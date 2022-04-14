package utils

import (
	"image"
	"image/color/palette"
	"image/gif"
	"image/png"
	"io"
)

func EncodePNG(w io.Writer) func(*image.RGBA) error {
	return func(drawableImg *image.RGBA) error {
		if err := png.Encode(w, drawableImg); err != nil {
			return err
		}
		return nil
	}
}

func EncodeGIF(w io.Writer, delay []int) func([]*image.Image) error {
	return func(imgs []*image.Image) error {
		palettedImgs := make([]*image.Paletted, len(imgs))
		// convert image.Image to image.Palette
		for i, img := range imgs {
			palettedImgs[i] = ImageToPaletted(*img, palette.Plan9)
		}
		processedGIF := gif.GIF{Image: palettedImgs, Delay: delay}
		if err := gif.EncodeAll(w, &processedGIF); err != nil {
			return err
		}
		return nil
	}
}
