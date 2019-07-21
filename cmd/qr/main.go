package main

import (
	"fmt"
	"os"

	"github.com/alxmsl/qrencode-go/qrencode"
	"github.com/spf13/cobra"
)

var (
	level int
	text  string

	rootCmd = &cobra.Command{
		Use:   "qr",
		Short: "This example shows how to use library to display QR code for different output modes",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
)

func main() {
	rootCmd.InheritedFlags().IntVar(&level, "level", int(qrencode.ECLevelH), "error correction level")
	rootCmd.InheritedFlags().StringVar(&text, "text", "", "text value should be encoded into QR code")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
