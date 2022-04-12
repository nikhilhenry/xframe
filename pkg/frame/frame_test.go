package frame_test

import (
	"bytes"
	approvals "github.com/approvals/go-approval-tests"
	bezel "github.com/nikhilhenry/xframe/internal/bezel"
	"github.com/nikhilhenry/xframe/pkg/frame"
	"image"
	"os"
	"testing"
)

// Returns a test image of the home screen from the iPhone-13 simulator
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

	var screenshot = getScreenTestImage(t)

	deviceBezel := bezel.New(bezel.Iphone13Pro)

	t.Run("it generates an image with the device deviceBezel", func(t *testing.T) {

		buf := bytes.Buffer{}
		err := frame.GenerateFrameWithBezel(&buf, *deviceBezel, screenshot)

		if err != nil {
			t.Fatal(err)
		}

		reader := bytes.NewReader(buf.Bytes())
		approvals.VerifyWithExtension(t, reader, ".png")
	})
}

func BenchmarkGenerate(b *testing.B) {

	screenshot := getScreenTestImage(b)
	deviceBezel := bezel.New(bezel.Iphone13Pro)

	buf := bytes.Buffer{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = frame.GenerateFrameWithBezel(&buf, *deviceBezel, screenshot)
	}
}
