package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "myrientui",
	Short: "A utility for interacting with the Myrient archives.",
}

func init() {
	rootCmd.AddCommand(browseCmd)
	rootCmd.AddCommand(indexCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
