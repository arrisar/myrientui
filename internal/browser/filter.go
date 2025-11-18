package browser

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

var FilterPromptStyle lipgloss.Style = lipgloss.NewStyle().
	Padding(1, 0, 0, 1).
	Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#ECFD65"})

func GetFilterInput() textinput.Model {
	model := textinput.New()
	model.Prompt = "Filter: "
	model.PromptStyle = FilterPromptStyle
	return model
}
