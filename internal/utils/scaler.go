package utils

import (
	"golang.org/x/image/draw"
	"image"
)

func EncodeWithScale(dim Dimension, encode func(rgba *image.RGBA) error) func(img *image.RGBA) error {
	scaledDstImage := image.NewRGBA(image.Rect(0, 0, dim.Width, dim.Height))
	return func(img *image.RGBA) error {
		draw.NearestNeighbor.Scale(scaledDstImage, scaledDstImage.Bounds(), img, img.Bounds(), draw.Over, nil)
		err := encode(scaledDstImage)
		if err != nil {
			return err
		}
		return nil
	}
}
