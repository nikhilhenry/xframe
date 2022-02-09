package frame

import (
	"bytes"
	"embed"
	"image"
	"image/draw"
	"image/png"
	"io"
)

//go:embed assets
var assetsFs embed.FS

// GenerateFrameWithBezel Generates an image with the screenshot embedded within a device bezel
func GenerateFrameWithBezel(w io.Writer, screenImage image.Image) error {
	//get embedded device image
	deviceImageFile, err := assetsFs.ReadFile("assets/iphone-13-pro.png")
	deviceImage, _, err := image.Decode(bytes.NewReader(deviceImageFile))
	if err != nil {
		return err
	}

	//get image bounds
	deviceImageBounds := deviceImage.Bounds()
	screenImageBounds := screenImage.Bounds()

	destinationPoint := getDestinationPoint(deviceImageBounds, screenImageBounds)

	r := image.Rectangle{Min: destinationPoint, Max: destinationPoint.Add(screenImageBounds.Size())}

	//make image drawable
	drawableScreenImage := image.NewRGBA(image.Rect(0, 0, deviceImageBounds.Dx(), deviceImageBounds.Dy()))
	draw.Draw(drawableScreenImage, r, screenImage, screenImageBounds.Min, draw.Src)

	//copy device bezel onto drawn image
	overRect := image.Rectangle{Min: deviceImageBounds.Min, Max: deviceImageBounds.Max}
	draw.Draw(drawableScreenImage, overRect, deviceImage, deviceImageBounds.Min, draw.Over)

	err = png.Encode(w, drawableScreenImage)
	if err != nil {
		return err
	}
	return nil
}

func getDestinationPoint(deviceBounds image.Rectangle, screenBounds image.Rectangle) image.Point {
	xCoordinates := deviceBounds.Max.X/2 - screenBounds.Max.X/2
	yCoordinates := deviceBounds.Max.Y/2 - screenBounds.Max.Y/2
	return image.Point{X: xCoordinates, Y: yCoordinates}
}
