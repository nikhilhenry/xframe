package utils

import "testing"

func Test_getFilePath(t *testing.T) {

	t.Run("it returns output with input image name when '.'", func(t *testing.T) {
		const imagePath = "./frame/testdata/screen-home.png"
		const outputPath = "."
		got := getFilePath(imagePath, outputPath)
		want := "./screen-home-framed.png"

		if got != want {
			t.Errorf("expected '%q' but got '%q'", want, got)
		}
	})

	t.Run("it returns output path with output image path", func(t *testing.T) {
		const imagePath = "./frame/testdata/screen-home.png"
		const outputPath = "./test.png"
		got := getFilePath(imagePath, outputPath)
		want := "./test.png"

		if got != want {
			t.Errorf("expected '%q' but got '%q'", want, got)
		}
	})

	t.Run("it appends .png to path if not specifcied", func(t *testing.T) {
		const imagePath = "./frame/testdata/screen-home.png"
		const outputPath = "./test"
		got := getFilePath(imagePath, outputPath)
		want := "./test.png"

		if got != want {
			t.Errorf("expected '%q' but got '%q'", want, got)
		}
	})

}
