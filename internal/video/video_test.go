package video

import (
	"io"
	"os"
	"testing"
)

func TestDecode(t *testing.T) {
	screenVideo := getTestVideo(t)
	err, _ := Decode(screenVideo)
	if err != nil {
		t.Error(t)
	}
}

func getTestVideo(t testing.TB) io.Reader {
	t.Helper()
	reader, err := os.Open("../../pkg/frame/testdata/screen-video.mp4")
	defer func(reader *os.File) {
		err := reader.Close()
		if err != nil {
			t.Error(err)
		}
	}(reader)
	if err != nil {
		t.Error(err)
	}

	return reader
}
