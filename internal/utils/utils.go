package utils

import (
	"image"
	"image/png"
	"io"
)

func ImageEncoderPNG(w io.Writer, drawableImg *image.RGBA) error {
	if err := png.Encode(w, drawableImg); err != nil {
		return err
	}
	return nil
}
