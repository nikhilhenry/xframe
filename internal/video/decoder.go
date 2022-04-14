package video

import (
	"bytes"
	"fmt"
	"image"
	"os"
	"os/exec"
	"path"
)

//"-f", "mp4",
//"-i", "pipe:0", // take stdin as input
//"-vf", "fps=64", // set fps
//"-c:v", "png", // output to png
//"-f", "image2pipe", // specify output
//"pipe:1", // output to stdin

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
	resultBuffer := bytes.NewBuffer(make([]byte, 50*1024*1024)) // pre allocate 5MiB buffer

	cmd.Stderr = os.Stderr    // bind log stream to stderr
	cmd.Stdout = resultBuffer // stdout result will be written here

	// open stdin pipe
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err, nil
	}
	// start process on another goroutine
	if err := cmd.Start(); err != nil {
		return err, nil
	}
	// pump video data to Stdin pipe
	//if _, err := stdin.Write(buf); err != nil {
	//	return err, nil
	//}
	// close stdin
	if err := stdin.Close(); err != nil {
		return err, nil
	}
	// wait for ffmpeg to finish
	if err := cmd.Wait(); err != nil {
		return err, nil
	}
	return nil, []image.Image{}
}
