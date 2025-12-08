package browser

import (
	"github.com/charmbracelet/lipgloss"
)

var BrowserStyle lipgloss.Style = lipgloss.NewStyle().
	Padding(1, 1, 1, 1)

/**
 * HEADER
 */

var HeaderStyle lipgloss.Style = lipgloss.NewStyle().
	Padding(0, 0, 0, 0)

var LogoStyle lipgloss.Style = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: "#9c3a14ff", Dark: "#a94118"}).
	Padding(0, 0, 1, 0).
	Bold(true)

var PathLabelStyle lipgloss.Style = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: "#c65ace", Dark: "#c65ace"}).
	Bold(true)

var PathValueStyle lipgloss.Style = lipgloss.NewStyle()

var FilterCursorStyle lipgloss.Style = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: "#c65ace", Dark: "#c65ace"})

var FilterPlaceholderStyle lipgloss.Style = lipgloss.NewStyle()

var FilterPromptStyle lipgloss.Style = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#ECFD65"}).
	Padding(0, 0, 0, 0).
	Bold(true)

var FilterValueStyle lipgloss.Style = lipgloss.NewStyle()

/**
 * LIST
 */

var ListStyle lipgloss.Style = lipgloss.NewStyle().
	Padding(1, 0, 0, 0)

var ListItemStyle lipgloss.Style = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: "#444844", Dark: "#BBBFBB"})

var ListCurrentStyle lipgloss.Style = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: "#115511", Dark: "#BBFFBB"})

var ListCountStyle lipgloss.Style = lipgloss.NewStyle().
	Padding(1, 0, 0, 0).
	Foreground(lipgloss.AdaptiveColor{Light: "#332233", Dark: "#998899"})

/**
 * FOOTER
 */

var FooterStyle lipgloss.Style = lipgloss.NewStyle().
	Padding(1, 0, 1, 0)

var KeysStyle lipgloss.Style = lipgloss.NewStyle()

var KeyValueStyle lipgloss.Style = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: "#336633", Dark: "#FFCC99"}).
	Bold(true)

var KeyTextStyle lipgloss.Style = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: "#332233", Dark: "#998899"})
