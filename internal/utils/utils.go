package utils

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"path"
	"strings"
)

func ImageEncoderPNG(w io.Writer) func(*image.RGBA) error {
	return func(drawableImg *image.RGBA) error {
		if err := png.Encode(w, drawableImg); err != nil {
			return err
		}
		return nil
	}
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
