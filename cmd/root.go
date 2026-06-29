package cmd

import (
	"context"

	"github.com/indium114/sorta/internal"

	"github.com/charmbracelet/fang"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var dir string
var dry bool

var rootCmd = &cobra.Command{
	Use:   "sorta [directory]",
	Short: "A simple CLI tool to sort files based on MIME type",
	Args:  cobra.MaximumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			dir = args[0]
		}

		if dir == "" {
			dir = "."
		}

		internal.RunSorter(dir, dry)
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&dry, "dry", "d", false, "dry run; don't actually move any files")
}

func Execute() {
	if err := fang.Execute(context.Background(), rootCmd); err != nil {
		log.Fatal(err)
	}
}
