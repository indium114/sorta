package cmd

import (
	"context"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sorta",
	Short: "A simple CLI tool to sort files based on MIME type",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	err := fang.Execute(context.Background(), rootCmd)
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
