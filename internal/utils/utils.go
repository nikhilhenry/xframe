package utils

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"path"
	"strconv"
	"strings"
)

type Dimension struct {
	Width  int
	Height int
}

func GetDimensionsFromFlag(flag string) (Dimension, error) {
	values := strings.Split(flag, "x")
	width, err := strconv.Atoi(values[0])
	height, err := strconv.Atoi(values[1])
	if err != nil {
		return Dimension{}, err
	}
	return Dimension{width, height}, nil
}

func GetFilePath(imageFilePath string, rawPath string) (filePath string) {
	if rawPath != "." {
		if strings.HasSuffix(rawPath, ".png") {
			filePath = rawPath
			return
		}
		filePath = rawPath + ".png"
		return
	}
	cleanedPath := path.Clean(rawPath)
	filePathElements := strings.SplitAfter(imageFilePath, "/")
	fmt.Println(filePathElements)
	fileNameWithExtension := filePathElements[len(filePathElements)-1]
	fileName := strings.Split(fileNameWithExtension, ".")[0]
	filePath = fmt.Sprintf("%s/%s-framed.png", cleanedPath, fileName)
	return
}
func ImageToPaletted(img image.Image, palette color.Palette) *image.Paletted {
	bounds := img.Bounds()
	palettedImage := image.NewPaletted(bounds, palette)
	draw.FloydSteinberg.Draw(palettedImage, bounds, img, image.Point{})
	return palettedImage
}
func PalettedToImage(img *image.Paletted) image.Image {
	bounds := img.Bounds()
	palettedImage := image.NewRGBA64(bounds)
	draw.FloydSteinberg.Draw(palettedImage, bounds, img, image.Point{})
	return palettedImage
}
