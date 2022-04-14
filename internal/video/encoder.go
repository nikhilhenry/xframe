package video

import (
	"fmt"
	"github.com/nikhilhenry/xframe/internal/utils"
	"image"
	"os"
	"path/filepath"
	"sync"
)

func Encode(videoPath string, imgs []image.Image) error {
	// make a temp dir to store encoded images
	dirName, err := os.MkdirTemp("", "xframe-frame-store")
	fmt.Println(dirName)
	//defer os.RemoveAll(dirName)
	if err != nil {
		return err
	}
	// encode all images to png
	var wg sync.WaitGroup
	for i, img := range imgs {
		wg.Add(1)
		go func(index int, rgba image.Image) {
			defer wg.Done()
			// make a new image
			outputPath, err := os.Create(filepath.Join(dirName, fmt.Sprintf("img-out-temp%v.png", index)))
			if err != nil {
				fmt.Println(err)
			}
			imageEncoder := utils.EncodePNG(outputPath)
			if err := imageEncoder(rgba); err != nil {
				fmt.Println(err)
			}
		}(i, img)
	}
	wg.Wait()
	return nil
}
