package main

import (
	"fmt"
	"os"

	"github.com/alxmsl/qrencode-go/output"
	"github.com/alxmsl/qrencode-go/qrencode"
	"github.com/spf13/cobra"
)

var (
	terminalSymbol string

	terminalBlackColor, terminalWhiteColor int

	terminalCmd = &cobra.Command{
		Use:   "terminal",
		Short: "Creates QR code for terminal output",
		Run: func(cmd *cobra.Command, args []string) {
			if !isValidColor(terminalBlackColor) || !isValidColor(terminalWhiteColor) {
				fmt.Println("Invalid terminal color")
				os.Exit(-1)
			}

			grid, err := qrencode.Encode(text, qrencode.ECLevel(level))
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}

			o := output.NewTerminalOutput(
				&output.TerminalSymbol{terminalBlackColor, terminalSymbol},
				&output.TerminalSymbol{terminalWhiteColor, terminalSymbol})
			o.Output(grid, os.Stdout)
		},
	}
)

func init() {
	terminalCmd.PersistentFlags().StringVar(&terminalSymbol, "symbol", "   ", "QR symbol representation for terminal")
	terminalCmd.PersistentFlags().IntVar(&terminalBlackColor, "black", output.TermBlack, "Color for black symbol")
	terminalCmd.PersistentFlags().IntVar(&terminalWhiteColor, "white", output.TermWhite, "Color for white symbol")
	rootCmd.AddCommand(terminalCmd)
}

func isValidColor(c int) bool {
	return c >= output.TermBlack && c <= output.TermWhite ||
		c >= output.TermBlackBright && c <= output.TermWhiteBright
}
