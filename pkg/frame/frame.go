package frame

import (
	"github.com/nikhilhenry/xframe/internal/bezel"
	"image"
	"image/draw"
)

// GenerateFrameWithBezel Generates an image with the screenshot embedded within a device bezel
func GenerateFrameWithBezel(encode func(rgba *image.RGBA) error, bezel bezel.Bezel, screenImage image.Image) error {

	//get image bounds
	deviceImageBounds := bezel.Bounds
	screenImageBounds := screenImage.Bounds()

	destinationPoint := getDestinationPoint(deviceImageBounds, screenImageBounds)

	r := image.Rectangle{Min: destinationPoint, Max: destinationPoint.Add(screenImageBounds.Size())}

	//make image drawable
	drawableScreenImage := image.NewRGBA(image.Rect(0, 0, deviceImageBounds.Dx(), deviceImageBounds.Dy()))
	draw.Draw(drawableScreenImage, r, screenImage, screenImageBounds.Min, draw.Src)

	//copy device bezel onto drawn image
	overRect := image.Rectangle{Min: deviceImageBounds.Min, Max: deviceImageBounds.Max}
	draw.Draw(drawableScreenImage, overRect, *bezel.Image, deviceImageBounds.Min, draw.Over)

	if err := encode(drawableScreenImage); err != nil {
		return err
	}
	return nil
}

func getDestinationPoint(deviceBounds image.Rectangle, screenBounds image.Rectangle) image.Point {
	xCoordinates := deviceBounds.Max.X/2 - screenBounds.Max.X/2
	yCoordinates := deviceBounds.Max.Y/2 - screenBounds.Max.Y/2
	return image.Point{X: xCoordinates, Y: yCoordinates}
}
