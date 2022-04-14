package bezel

import (
	"bytes"
	"embed"
	"fmt"
	"image"
	"log"
)

// Iphone enum of supported bezels
const (
	Iphone13Pro string = "iphone-13-pro.png"
)

//go:embed assets
var assetsFs embed.FS

// Bezel holds the Name of a iphone device bezel
// Default is iphone-13-pro
// @todo improve performance by storing pointer to the file and bounds instead of computing on each call
type Bezel struct {
	Name string
}

// Image returns a pointer to the image file of the bezel
func (b Bezel) Image() *image.Image {
	switch b.Name {
	case Iphone13Pro:
		err, img := load(Iphone13Pro)
		if err != nil {
			log.Fatalln("failed to load image")
		}
		return img
	default:
		err, img := load(Iphone13Pro)
		if err != nil {
			log.Fatalln("failed to load image")
		}
		return img
	}
}

// Bounds returns the image bound of the bezel
func (b Bezel) Bounds() image.Rectangle {
	img := *b.Image()
	return img.Bounds()
}

func load(name string) (error, *image.Image) {
	deviceImageFile, err := assetsFs.ReadFile(fmt.Sprintf("assets/%v", name))
	deviceImage, _, err := image.Decode(bytes.NewReader(deviceImageFile))
	if err != nil {
		return err, nil
	}
	return nil, &deviceImage
}
