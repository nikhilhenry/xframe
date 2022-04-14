package video

import (
	"bytes"
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

func getTestVideo(t testing.TB) []byte {
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
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	return buf.Bytes()
}
