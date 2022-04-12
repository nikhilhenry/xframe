package utils

import (
	"golang.org/x/image/draw"
	"image"
)

func EncodeWithScale(dim Dimension, encode func(rgba *image.RGBA) error) func(img *image.RGBA) error {
	// Do not scale if either Dimension value is 0
	if dim.Width == 0 || dim.Height == 0 {
		return func(img *image.RGBA) error {
			if err := encode(img); err != nil {
				return err
			}
			return nil
		}
	}
	scaledDstImage := image.NewRGBA(image.Rect(0, 0, dim.Width, dim.Height))
	return func(img *image.RGBA) error {
		draw.ApproxBiLinear.Scale(scaledDstImage, scaledDstImage.Bounds(), img, img.Bounds(), draw.Over, nil)
		err := encode(scaledDstImage)
		if err != nil {
			return err
		}
		return nil
	}
}
