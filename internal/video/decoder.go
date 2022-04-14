package video

import (
	"fmt"
	"image"
	"os"
	"os/exec"
	"path"
)

func Decode(filePath string) (error, []image.Image) {

	// create a temporary directory to store images
	dirName, err := os.MkdirTemp("", "xframe-frame-store")
	defer os.RemoveAll(dirName)
	if err != nil {
		return err, nil
	}
	fmt.Println(dirName)
	outputPath := fmt.Sprintf("%v/img-out-temp%%d.png", dirName)
	cmd := exec.Command("ffmpeg", "-y", // yes to all
		"-i", path.Clean(filePath), // take stdin as input
		"-vf", "fps=64", // set fps
		outputPath, // output to stdin
	)

	cmd.Stderr = os.Stderr // bind log stream to stderr

	// start process on another goroutine
	if err := cmd.Start(); err != nil {
		return err, nil
	}
	// wait for ffmpeg to finish
	if err := cmd.Wait(); err != nil {
		return err, nil
	}
	return nil, []image.Image{}
}
