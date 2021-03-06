package frame_test

import (
	"bytes"
	approvals "github.com/approvals/go-approval-tests"
	"github.com/nikhilhenry/xframe/internal/bezel"
	"github.com/nikhilhenry/xframe/internal/utils"
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

	deviceBezel := bezel.Bezel{Name: bezel.Iphone13Pro}

	t.Run("it generates an image with the device bezel", func(t *testing.T) {

		buf := bytes.Buffer{}
		err := frame.Generate(utils.EncodePNG(&buf), deviceBezel, screenshot)

		if err != nil {
			t.Fatal(err)
		}

		reader := bytes.NewReader(buf.Bytes())
		approvals.VerifyWithExtension(t, reader, ".png")
	})
	t.Run("it generates an image scaled to the provided dimension", func(t *testing.T) {
		dim := utils.Dimension{Width: 685, Height: 1356}
		buf := bytes.Buffer{}
		err := frame.Generate(utils.EncodeWithScale(dim, utils.EncodePNG(&buf)), deviceBezel, screenshot)

		if err != nil {
			t.Fatal(err)
		}

		reader := bytes.NewReader(buf.Bytes())
		approvals.VerifyWithExtension(t, reader, ".png")
	})

	t.Run("it generates an image not scaled when no dimension is provided", func(t *testing.T) {
		dim := utils.Dimension{Width: 0, Height: 0}
		buf := bytes.Buffer{}
		err := frame.Generate(utils.EncodeWithScale(dim, utils.EncodePNG(&buf)), deviceBezel, screenshot)

		if err != nil {
			t.Fatal(err)
		}

		reader := bytes.NewReader(buf.Bytes())
		approvals.VerifyWithExtension(t, reader, ".png")
	})
}

func BenchmarkGenerate(b *testing.B) {

	screenshot := getScreenTestImage(b)
	deviceBezel := bezel.Bezel{Name: bezel.Iphone13Pro}

	buf := bytes.Buffer{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = frame.Generate(utils.EncodePNG(&buf), deviceBezel, screenshot)
	}
}
