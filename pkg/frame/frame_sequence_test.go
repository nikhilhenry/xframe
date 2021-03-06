package frame_test

import (
	"bytes"
	approvals "github.com/approvals/go-approval-tests"
	"github.com/nikhilhenry/xframe/internal/bezel"
	"github.com/nikhilhenry/xframe/internal/utils"
	"github.com/nikhilhenry/xframe/internal/video"
	"github.com/nikhilhenry/xframe/pkg/frame"
	"image"
	_ "image/gif"
	"io"
	"os"
	"testing"
)

func TestGenerateSequence(t *testing.T) {
	approvals.UseFolder("testdata")
	deviceBezel := bezel.Bezel{Name: bezel.Iphone13Pro}
	t.Run("it generates a gif with the device bezel", func(t *testing.T) {
		var screenshot = getScreenTestGIF(t)
		imgs := utils.DecodeGIF(screenshot)

		buf := bytes.Buffer{}
		err := frame.GenerateSequence(utils.EncodeGIF(&buf, screenshot.Delay), deviceBezel, imgs)

		if err != nil {
			t.Fatal(err)
		}

		reader := bytes.NewReader(buf.Bytes())
		approvals.VerifyWithExtension(t, reader, ".gif")
	})

	t.Run("it generates a video with the device bezel", func(t *testing.T) {

		imgs := getTestVideo(t)
		err := frame.GenerateSequence(video.EncodeImgs("./testdata/screen-video-frame.mp4"), deviceBezel, imgs)

		if err != nil {
			t.Fatal(err)
		}

		reader := readVideoFile(t)
		approvals.VerifyWithExtension(t, reader, ".mp4")
	})
}

func getTestVideo(t testing.TB) []image.Image {
	t.Helper()
	err, imgs := video.Decode("./testdata/screen-video.mp4")
	if err != nil {
		t.Fatal(err)
	}
	return imgs
}
func readVideoFile(t testing.TB) io.Reader {
	reader, err := os.Open("./testdata/screen-video-frame.mp4")
	if err != nil {
		t.Fatal(err)
	}
	return reader
}
