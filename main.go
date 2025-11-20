package main

import (
	"fmt"
	"os"

	"github.com/arrisar/myrientui/internal/app"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	a := app.New()
	p := tea.NewProgram(a, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
