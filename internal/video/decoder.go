package video

import (
	"bytes"
	"image"
	"io"
	"os"
	"os/exec"
)

func Decode(reader io.Reader) (error, []image.Image) {

	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	cmd := exec.Command("ffmpeg", "y", // yes to all
		"-i", "pipe:0", // take stdin as input
		"-vf", "fps=64", // set fps
		"pipe:1", // output to stdin
	)
	resultBuffer := bytes.NewBuffer(make([]byte, 5*1024*1024)) // pre allocate 5MiB buffer

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
	if _, err := stdin.Write(buf.Bytes()); err != nil {
		return err, nil
	}
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
