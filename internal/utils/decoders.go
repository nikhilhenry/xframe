package utils

import (
	"image"
	"image/gif"
)

func DecodeGIF(gif gif.GIF) []*image.Image {
	imgs := gif.Image
	processedImages := make([]*image.Image, len(imgs))
	for i, img := range imgs {
		processedImage := PalettedToImage(img)
		processedImages[i] = &processedImage
	}
	return processedImages
}
