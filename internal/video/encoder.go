package video

import (
	"fmt"
	"github.com/nikhilhenry/xframe/internal/utils"
	"image"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

var encodeFileTemplate = "img-out-temp%v.png"

func Encode(videoPath string, imgs []image.Image) error {
	// make a temp dir to store encoded images
	dirName, err := os.MkdirTemp("", "xframe-frame-store")
	fmt.Println(dirName)
	defer os.RemoveAll(dirName)
	if err != nil {
		return err
	}
	// encode all images to png
	var wg sync.WaitGroup
	for i, img := range imgs {
		wg.Add(1)
		go func(index int, rgba image.Image) {
			defer wg.Done()
			// make a new image
			outputPath, err := os.Create(filepath.Join(dirName, fmt.Sprintf(encodeFileTemplate, index)))
			if err != nil {
				fmt.Println(err)
			}
			imageEncoder := utils.EncodePNG(outputPath)
			if err := imageEncoder(rgba); err != nil {
				fmt.Println(err)
			}
		}(i, img)
	}
	wg.Wait()
	// now pass these images to ffmpeg and encode
	cmd := exec.Command("ffmpeg", "-y",
		"-framerate", "64",
		"-f", "image2",
		"-i", filepath.Join(dirName, fmt.Sprintf(encodeFileTemplate, "%d")),
		"-vcodec", "libx264",
		"-pix_fmt", "yuv420p",
		filepath.Join(videoPath),
	)
	cmd.Stderr = os.Stderr // bind log stream to stderr

	// start process on another goroutine
	if err := cmd.Start(); err != nil {
		return err
	}
	// wait for ffmpeg to finish
	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}
