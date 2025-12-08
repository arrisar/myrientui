package cmd

import (
	"fmt"
	"os"

	"github.com/arrisar/myrientui/internal/app"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var browseCmd = &cobra.Command{
	Use:   "browse",
	Short: "Launch the MyrienTUI browser interface.",
	Run: func(cmd *cobra.Command, args []string) {
		a := app.New()
		p := tea.NewProgram(a, tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	},
}
