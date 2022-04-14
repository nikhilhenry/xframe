package frame_test

import (
	"bytes"
	approvals "github.com/approvals/go-approval-tests"
	"github.com/nikhilhenry/xframe/internal/bezel"
	"github.com/nikhilhenry/xframe/internal/utils"
	"github.com/nikhilhenry/xframe/pkg/frame"
	_ "image/gif"
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
}
