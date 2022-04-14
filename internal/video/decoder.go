package video

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

type result struct {
	index int
	image image.Image
}

func Decode(videoPath string) (error, []image.Image) {

	// create a temporary directory to store images
	dirName, err := os.MkdirTemp("", "xframe-frame-store")
	defer os.RemoveAll(dirName)
	if err != nil {
		return err, nil
	}
	fmt.Println(dirName)
	outputPath := fmt.Sprintf("%v/img-out-temp%%d.png", dirName)
	cmd := exec.Command("ffmpeg", "-y", // yes to all
		"-i", path.Clean(videoPath), // take stdin as input
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

	resultsChannel := make(chan result)

	// read all files created by ffmpeg
	files, err := os.ReadDir(dirName)
	if err != nil {
		return err, nil
	}
	for i, file := range files {
		go func(index int, fileName string) error {
			//	read the file
			reader, err := os.Open(fileName)
			defer reader.Close()
			if err != nil {
				fmt.Println(err)
				return err
			}
			img, _, err := image.Decode(reader)
			if err != nil {
				fmt.Println(err)
				return err
			}
			resultsChannel <- result{
				index: index,
				image: img,
			}
			return nil
		}(i, filepath.Join(dirName, file.Name()))
	}
	imgs := make([]image.Image, len(files))
	for i := 0; i < len(files); i++ {
		r := <-resultsChannel
		imgs[r.index] = r.image
	}
	return nil, imgs
}
