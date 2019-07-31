package main

import (
	"fmt"
	"os"

	"github.com/alxmsl/qrencode-go/output"
	"github.com/alxmsl/qrencode-go/qrencode"
	"github.com/spf13/cobra"
)

const (
	processorRectangle = "rectangle"
	processorCircle    = "circle"
)

var (
	vectorProcessor string
	vectorBlock     int

	vectorCmd = &cobra.Command{
		Use:   "vector",
		Short: "Creates QR code for vector output",
		Run: func(cmd *cobra.Command, args []string) {
			var o output.Interface
			switch vectorProcessor {
			case processorCircle:
				o = output.NewVectorOutput(vectorBlock, 0, 0, output.DefaultCircleProcessor, nil)
			case processorRectangle:
				o = output.NewVectorOutput(vectorBlock, 0, 0, output.DefaultRectangleProcessor, nil)
			default:
				fmt.Println("Invalid vector processor")
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
	vectorCmd.PersistentFlags().StringVar(&vectorProcessor, "processor", processorRectangle, "Vector processor: circle, rectangle")
	vectorCmd.PersistentFlags().IntVar(&vectorBlock, "block", 4, "Vector block size")
	rootCmd.AddCommand(vectorCmd)
}
