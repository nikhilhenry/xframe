package frame_test

import (
	"bytes"
	approvals "github.com/approvals/go-approval-tests"
	"github.com/nikhilhenry/x-frame/frame"
	"image/gif"
	_ "image/gif"
	"os"
	"testing"
)

func TestGenerateGif(t *testing.T) {
	approvals.UseFolder("testdata")

	t.Run("it generates a gif with the device bezel", func(t *testing.T) {
		var screenshot = getScreenTestGIF(t)

		buf := bytes.Buffer{}
		err := frame.GenerateFrameWithBezelGIF(&buf, screenshot)

		if err != nil {
			t.Fatal(err)
		}

		reader := bytes.NewReader(buf.Bytes())
		approvals.VerifyWithExtension(t, reader, ".gif")
	})
}

func BenchmarkGenerateGif(b *testing.B) {

	var screenshot = getScreenTestGIF(b)

	buf := bytes.Buffer{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = frame.GenerateFrameWithBezelGIF(&buf, screenshot)
	}
}

// Returns a test GIF of the home screen from the iPhone-13 simulator
func getScreenTestGIF(t testing.TB) gif.GIF {
	t.Helper()
	reader, err := os.Open("testdata/screen-home.gif")
	defer func(reader *os.File) {
		err := reader.Close()
		if err != nil {
			t.Error(err)
		}
	}(reader)
	if err != nil {
		t.Error(err)
	}
	screenShotImage, err := gif.DecodeAll(reader)
	if err != nil {
		t.Error("unable to decode gif file")
	}
	return *screenShotImage
}
