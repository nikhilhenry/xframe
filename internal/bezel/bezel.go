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

// Bezel holds the Image, Bounds and Name of a iphone device bezel
type Bezel struct {
	Image  *image.Image
	Bounds image.Rectangle
	Name   string
}

// New returns a Bezel for a given device name.
// Default device is Iphone 13 Pro
func New(name string) *Bezel {
	switch name {
	case Iphone13Pro:
		err, img := load(Iphone13Pro)
		if err != nil {
			log.Fatalln("failed to load image")
		}
		// todo write better dereferencing
		bounds := *img
		return &Bezel{Image: img, Bounds: bounds.Bounds(), Name: name}
	default:
		return &Bezel{}
	}
}

func load(name string) (error, *image.Image) {
	deviceImageFile, err := assetsFs.ReadFile(fmt.Sprintf("assets/%v", name))
	deviceImage, _, err := image.Decode(bytes.NewReader(deviceImageFile))
	if err != nil {
		return err, nil
	}
	return nil, &deviceImage
}
