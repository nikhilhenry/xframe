package video

import (
	"github.com/nikhilhenry/xframe/internal/utils"
	"image/gif"
	"os"
	"testing"
)

func TestDecode(t *testing.T) {
	//screenVideo := getTestVideo(t)
	err, _ := Decode("../../pkg/frame/testdata/screen-video.mp4")
	if err != nil {
		t.Error(t)
	}
}

func TestEncode(t *testing.T) {
	var screenshot = getScreenTestGIF(t)
	imgs := utils.DecodeGIF(screenshot)
	err := Encode("./output/video-framed.mp4", imgs)
	if err != nil {
		t.Error(err)
	}
}
func getScreenTestGIF(t testing.TB) gif.GIF {
	t.Helper()
	reader, err := os.Open("../../pkg/frame/testdata/screen-home.gif")
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
