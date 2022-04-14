package video

import (
	"testing"
)

func TestDecode(t *testing.T) {
	//screenVideo := getTestVideo(t)
	err, _ := Decode("../../pkg/frame/testdata/screen-video.mp4")
	if err != nil {
		t.Error(t)
	}
}
