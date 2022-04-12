// Package cmd /*
package cmd

import (
	"bytes"
	"fmt"
	"github.com/nikhilhenry/xframe/internal/bezel"
	"github.com/nikhilhenry/xframe/internal/utils"
	"github.com/nikhilhenry/xframe/pkg/frame"
	"github.com/spf13/cobra"
	"image"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: "0.0.0",
	Use:     "xframe <path to simulator screenshot> <output path>",
	Example: "xframe screenshot.png .",
	Short:   "Generates screenshots with IOS device bezels overlay",
	Long:    `A CLI tool to draw device bezels on IOS screenshots from the Xcode simulator`,
	Args:    cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		// get dimension flag
		dimFlag, _ := cmd.Flags().GetString("dimension")
		dim, _ := utils.GetDimensionsFromFlag(dimFlag)
		fmt.Println(dim)
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
		outputPath := utils.GetFilePath(args[0], args[1])
		outputImage, fileErr := os.Create(outputPath)
		if fileErr != nil {
			return fileErr
		}
		defer func(outputImage *os.File) {
			err := outputImage.Close()
			if err != nil {

			}
		}(outputImage)

		deviceBezel := bezel.Bezel{Name: bezel.Iphone13Pro}
		err = frame.Generate(utils.EncodeWithScale(dim, utils.ImageEncoderPNG(outputImage)), deviceBezel, screenShotImage)
		if err != nil {
			return err
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.XFrame.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringP("dimension", "d", "0x0", "Dimension for output image. Example '1370x2712'")
}
