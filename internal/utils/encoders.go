package utils

import (
	"image"
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

func EncodeGIF(w io.Writer, delay []int) func([]*image.Paletted) error {
	return func(imgs []*image.Paletted) error {
		processedGIF := gif.GIF{Image: imgs, Delay: delay}
		if err := gif.EncodeAll(w, &processedGIF); err != nil {
			return err
		}
		return nil
	}
}
