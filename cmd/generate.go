// Package cmd /*
package cmd

import (
	"bytes"
	"fmt"
	"github.com/nikhilhenry/X-Frame/frame"
	"github.com/spf13/cobra"
	"image"
	"os"
	"path"
	"strings"
)

// generateCmd represents the frame command
var generateCmd = &cobra.Command{
	Use:   "frame",
	Short: "generates an output image with the screenshot over a device bezel",
	Long: `frame an output image with the screenshot over a desired simulator screenshot
using official Apple device bezels.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//read image
		file, err := os.ReadFile(args[0])
		if err != nil {
			return err
		}
		screenShotImage, _, err := image.Decode(bytes.NewReader(file))
		if err != nil {
			return err
		}
		//create output image
		outputPath := path.Clean(args[1])
		filePath := strings.SplitAfter(args[0], "/")
		fileNameWithExtension := filePath[len(filePath)-1]
		fmt.Println(fileNameWithExtension)
		fileName := strings.Split(fileNameWithExtension, ".")[0]
		outputImage, fileErr := os.Create(fmt.Sprintf("%s/%s-framed.png", outputPath, fileName))
		if fileErr != nil {
			return fileErr
		}
		defer func(outputImage *os.File) {
			err := outputImage.Close()
			if err != nil {

			}
		}(outputImage)

		err = frame.GenerateFrameWithBezel(outputImage, screenShotImage)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
