package frame_test

import (
	"bytes"
	approvals "github.com/approvals/go-approval-tests"
	"github.com/nikhilhenry/X-Frame/frame"
	"image"
	"os"
	"testing"
)

func getScreenTestImage(t testing.TB) image.Image {
	t.Helper()
	reader, err := os.Open("testdata/screen-home.png")
	defer func(reader *os.File) {
		err := reader.Close()
		if err != nil {
			t.Error(err)
		}
	}(reader)
	if err != nil {
		t.Error(err)
	}
	screenShotImage, _, err := image.Decode(reader)
	if err != nil {
		t.Error("unable to decode image file")
	}
	return screenShotImage
}

func TestGenerate(t *testing.T) {
	approvals.UseFolder("testdata")

	t.Run("it generates an image with the device bezel", func(t *testing.T) {
		var screenshot = getScreenTestImage(t)

		buf := bytes.Buffer{}
		err := frame.Generate(&buf, screenshot)

		if err != nil {
			t.Fatal(err)
		}

		reader := bytes.NewReader(buf.Bytes())
		approvals.VerifyWithExtension(t, reader, ".png")
	})
}

func BenchmarkGenerate(b *testing.B) {

	var screenshot = getScreenTestImage(b)

	buf := bytes.Buffer{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = frame.Generate(&buf, screenshot)
	}
}
