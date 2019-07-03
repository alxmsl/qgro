package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alxmsl/qrencode-go/output"
	"github.com/alxmsl/qrencode-go/qrencode"
	"github.com/spf13/cobra"
)

const (
	OutputPng      = "png"
	OutputSvgCirc  = "svg+circ"
	OutputSvgRect  = "svg+rect"
	OutputTerminal = "terminal"
)

var (
	error int
	image string
	mode  string
	text  string

	rootCmd = &cobra.Command{
		Use:   "example",
		Short: "This example shows how to use library to display QR code for different output modes",
		Run: func(cmd *cobra.Command, args []string) {
			grid, err := qrencode.Encode(text, qrencode.ECLevel(error))
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}

			var i *output.VectorImage
			if image != "" {
				i = &output.VectorImage{
					Filename: image,
					Size:     (grid.Width()*8 + 8*8) / 4,
				}
			}

			var o output.Interface
			switch mode {
			case OutputPng:
				o = output.NewImageOutput(output.Png, 8, 16, nil, nil)
			case OutputSvgCirc:
				o = output.NewVectorOutput(8, 0, 0, output.CircleProcessor, i)
			case OutputSvgRect:
				o = output.NewVectorOutput(8, 0, 0, output.RectangleProcessor, i)
			case OutputTerminal:
				o = output.NewTerminalOutput(
					&output.TerminalSymbol{output.TermBlack, "   "},
					&output.TerminalSymbol{output.TermWhite, "   "})
			default:
				fmt.Printf("value %s is not correct mode", mode)
				os.Exit(-1)
			}
			o.Output(grid, os.Stdout)
		},
	}

	rvv = []string{
		OutputTerminal,
		OutputPng,
		OutputSvgRect,
	}
)

func main() {
	rootCmd.PersistentFlags().IntVar(&error, "error", int(qrencode.ECLevelH), "error correction level")
	rootCmd.PersistentFlags().StringVar(&image, "image", "", "image to mix-in the QR code for SVG mode")
	rootCmd.PersistentFlags().StringVar(&mode, "mode", OutputTerminal, fmt.Sprintf("output mode: %s", strings.Join(rvv, ", ")))
	rootCmd.PersistentFlags().StringVar(&text, "text", "", "text value should be encoded into QR code")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
