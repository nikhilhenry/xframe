package utils

import (
	"image"
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
