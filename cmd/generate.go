// Package cmd /*
package cmd

import (
	"bytes"
	"fmt"
	"github.com/nikhilhenry/xframe/pkg/frame"
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
	Args: cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		//@todo check if image is gif
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
		outputPath := getFilePath(args[0], args[1])
		outputImage, fileErr := os.Create(outputPath)
		if fileErr != nil {
			return fileErr
		}
		defer func(outputImage *os.File) {
			err := outputImage.Close()
			if err != nil {

			}
		}(outputImage)

		err = frame.Generate(outputImage, screenShotImage)
		if err != nil {
			return err
		}
		return nil
	},
}

func getFilePath(imageFilePath string, rawPath string) (filePath string) {
	if rawPath != "." {
		if strings.HasSuffix(rawPath, ".png") {
			filePath = rawPath
			return
		}
		filePath = rawPath + ".png"
		return
	}
	cleanedPath := path.Clean(rawPath)
	filePathElements := strings.SplitAfter(imageFilePath, "/")
	fmt.Println(filePathElements)
	fileNameWithExtension := filePathElements[len(filePathElements)-1]
	fileName := strings.Split(fileNameWithExtension, ".")[0]
	filePath = fmt.Sprintf("%s/%s-framed.png", cleanedPath, fileName)
	return
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
