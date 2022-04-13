package utils

import (
	"fmt"
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
