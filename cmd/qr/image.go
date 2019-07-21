package main

import (
	"fmt"
	"image/color"
	"os"

	"github.com/alxmsl/qrencode-go/output"
	"github.com/alxmsl/qrencode-go/qrencode"
	"github.com/spf13/cobra"
)

const (
	formatGif  = "gif"
	formatJpeg = "jpeg"
	formatPng  = "png"
)

var (
	imageFormat string

	imageBlock, imageMargin          int
	imageBlackColor, imageWhiteColor uint16

	imageCmd = &cobra.Command{
		Use:   "image",
		Short: "Creates QR code for image output",
		Run: func(cmd *cobra.Command, args []string) {
			var o output.Interface
			switch imageFormat {
			case formatGif:
				o = output.NewGifOutput(imageBlock, imageMargin, color.Gray16{imageBlackColor}, color.Gray16{imageWhiteColor}, nil)
			case formatJpeg:
				o = output.NewJpegOutput(imageBlock, imageMargin, color.Gray16{imageBlackColor}, color.Gray16{imageWhiteColor}, nil)
			case formatPng:
				o = output.NewPngOutput(imageBlock, imageMargin, color.Gray16{imageBlackColor}, color.Gray16{imageWhiteColor})
			default:
				fmt.Println("Invalid image format")
				os.Exit(-1)
			}

			grid, err := qrencode.Encode(text, qrencode.ECLevel(level))
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
			o.Output(grid, os.Stdout)
		},
	}
)

func init() {
	imageCmd.PersistentFlags().StringVar(&imageFormat, "format", formatPng, "Image format: gif, jpeg, png")
	imageCmd.PersistentFlags().IntVar(&imageBlock, "block", 4, "Image block size, px")
	imageCmd.PersistentFlags().IntVar(&imageMargin, "margin", 8, "Image margin size, px")
	imageCmd.PersistentFlags().Uint16Var(&imageBlackColor, "black", 0, "Color for black symbol")
	imageCmd.PersistentFlags().Uint16Var(&imageWhiteColor, "white", 0xffff, "Color for white symbol")
	rootCmd.AddCommand(imageCmd)
}
