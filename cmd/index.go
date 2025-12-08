package cmd

import (
	"github.com/arrisar/myrientui/internal/config"
	"github.com/arrisar/myrientui/internal/scraper"
	"github.com/spf13/cobra"
)

var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "Run a recursive index of all available files/folders.",
	Run: func(cmd *cobra.Command, args []string) {
		c := config.New()
		s := scraper.New(c.Data.Storage.CacheDir)
		s.IndexAll()
	},
}
