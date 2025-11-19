package browser

import (
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type Footer struct {
	enabled bool
	keys    []Key

	loading bool
	spinner spinner.Model

	filterEnabled bool
	filterActive  bool
	filterApplied bool

	height int
	width  int
}

func NewFooter() Footer {
	f := Footer{}

	f.enabled = true
	f.keys = []Key{}

	f.loading = true
	f.spinner = spinner.New()

	f.filterEnabled = true
	f.filterActive = false
	f.filterApplied = false
	f.RebuildKeys()

	f.height = 0
	f.width = 0

	return f
}

/**
 * GETTERS
 */

func (f Footer) IsEnabled() bool {
	return f.enabled
}

/**
 * SETTERS
 */

func (f *Footer) SetEnabled(v bool) *Footer {
	f.enabled = v
	return f
}

func (f *Footer) RebuildKeys() *Footer {
	var next []Key

	switch true {
	case f.filterActive:
		next = []Key{
			{"enter", "apply"},
			{"esc", "clear"},
		}
	default:
		next = []Key{
			{"enter", "open"},
			{"space", "select"},
			{"f", "filter"},
		}
	}

	f.keys = next
	return f
}

/**
 * MODEL
 */

func (f Footer) Init() tea.Cmd {
	return tea.Batch(f.spinner.Tick,
		func() tea.Msg {
			f.RebuildKeys()
			return nil
		})
}

func (f Footer) View() (content string) {
	if !f.enabled {
		return
	}

	style := FooterStyle.Height(f.height).Width(f.width)
	if f.loading {
		return style.Render(f.spinner.View() + " Loading options...")
	}

	var keys []string
	for _, key := range f.keys {
		keys = append(keys, key.Render())
	}

	content = strings.Join(keys, " • ")
	return style.Render(content)
}

func (f Footer) Update(msg tea.Msg) (Footer, tea.Cmd) {
	var c, cmd tea.Cmd

	f, cmd = f.handleUpdateMsg(msg)
	c = tea.Batch(c, cmd)

	f.spinner, cmd = f.spinner.Update(msg)
	c = tea.Batch(c, cmd)

	return f, c
}

/**
 * HANDLERS
 */

func (f Footer) handleUpdateMsg(msg tea.Msg) (Footer, tea.Cmd) {
	switch msg := msg.(type) {
	case BrowserResizeMsg:
		return f.handleBrowserResizeMsg(msg)
	case FilterStateMsg:
		return f.handleFilterStateMsg(msg)
	case OptionsChangedMsg:
		return f.handleOptionsChangedMsg(msg)
	case OptionsLoadingMsg:
		return f.handleOptionsLoadingMsg(msg)
	default:
		return f, nil
	}
}

func (f Footer) handleBrowserResizeMsg(msg BrowserResizeMsg) (Footer, tea.Cmd) {
	f.height = msg.FooterHeight
	f.width = msg.FooterWidth
	return f, nil
}

func (f Footer) handleFilterStateMsg(msg FilterStateMsg) (Footer, tea.Cmd) {
	f.filterEnabled = msg.Enabled
	f.filterActive = msg.Active
	f.filterApplied = msg.Applied
	f.RebuildKeys()
	return f, nil
}

func (f Footer) handleOptionsChangedMsg(OptionsChangedMsg) (Footer, tea.Cmd) {
	f.loading = false
	return f, nil
}

func (f Footer) handleOptionsLoadingMsg(OptionsLoadingMsg) (Footer, tea.Cmd) {
	f.loading = true
	return f, nil
}
