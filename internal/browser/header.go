package browser

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Header struct {
	enabled   bool
	pathLabel string
	pathValue string

	Filter        textinput.Model
	filterPrompt  string
	filterValue   string
	filterActive  bool
	filterApplied bool
	filterEnabled bool

	height int
	width  int
}

func NewHeader() Header {
	h := Header{}

	h.enabled = true
	h.pathLabel = "Path:"
	h.pathValue = "/"

	h.Filter = textinput.New()
	h.filterPrompt = "Filter:"
	h.filterValue = ""
	h.filterActive = false
	h.filterApplied = false
	h.filterEnabled = true

	h.height = 0
	h.width = 0

	return h
}

/**
 * SETTERS
 */

func (h *Header) SetEnabled(v bool) *Header {
	h.enabled = v
	return h
}

/**
 * MODEL
 */

func (h Header) Init() tea.Cmd {
	return nil
}

func (h Header) View() (content string) {
	switch true {
	case !h.enabled:
		return
	case h.filterEnabled && (h.filterActive || h.filterApplied):
		h.Filter.PromptStyle = FilterPromptStyle
		h.Filter.Prompt = h.filterPrompt
		content = h.Filter.View()
	default:
		content = fmt.Sprint(
			PathLabelStyle.Render(h.pathLabel),
			" ",
			PathValueStyle.Render(h.pathValue),
		)
	}

	return HeaderStyle.Render(content)
}

func (h Header) Update(msg tea.Msg) (Header, tea.Cmd) {
	var c, cmd tea.Cmd

	h, cmd = h.handleUpdateMsg(msg)
	c = tea.Batch(c, cmd)

	if h.filterEnabled && c == nil {
		h.Filter, cmd = h.Filter.Update(msg)
		c = tea.Batch(c, cmd)
	}

	if h.filterValue != h.Filter.Value() {
		h.filterValue = h.Filter.Value()
		c = tea.Batch(c, func() tea.Msg {
			return FilterStateMsg{
				Enabled: h.filterEnabled,
				Active:  h.filterActive,
				Applied: h.filterApplied,
				Value:   h.filterValue,
			}
		})
	}

	return h, c
}

/**
 * HANDLERS
 */

func (h Header) handleUpdateMsg(msg tea.Msg) (Header, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return h.handleKeyMsg(msg)
	case BrowserResizeMsg:
		return h.handleBrowserResizeMsg(msg)
	case OptionSelectedMsg:
		return h.handleOptionSelectedMsg(msg)
	default:
		return h, nil
	}
}

func (h Header) handleKeyMsg(msg tea.KeyMsg) (Header, tea.Cmd) {
	var cmd tea.Cmd

	// active filter controls
	if h.filterEnabled && h.filterActive {
		switch msg.String() {
		case "enter":
			h.Filter.Blur()
			h.filterActive = false
			h.filterApplied = len(h.Filter.Value()) > 0
			h.filterValue = h.Filter.Value()
			cmd = func() tea.Msg {
				return FilterStateMsg{
					Enabled: h.filterEnabled,
					Active:  h.filterActive,
					Applied: h.filterApplied,
					Value:   h.filterValue,
				}
			}
		case "esc":
			h.Filter.SetValue("")
			h.Filter.Blur()
			h.filterActive = false
			h.filterApplied = false
			h.filterValue = ""
			cmd = func() tea.Msg {
				return FilterStateMsg{
					Enabled: h.filterEnabled,
					Active:  h.filterActive,
					Applied: h.filterApplied,
					Value:   h.filterValue,
				}
			}
		}
	}

	// inactive filter controls
	if h.filterEnabled && !h.filterActive {
		switch msg.String() {
		case "f":
			h.Filter.Focus()
			h.filterActive = true
			h.filterApplied = len(h.Filter.Value()) > 0
			h.filterValue = h.Filter.Value()
			cmd = func() tea.Msg {
				return FilterStateMsg{
					Enabled: h.filterEnabled,
					Active:  h.filterActive,
					Applied: h.filterApplied,
					Value:   h.filterValue,
				}
			}
		}
	}

	return h, cmd
}

func (h Header) handleBrowserResizeMsg(msg BrowserResizeMsg) (Header, tea.Cmd) {
	h.height = msg.HeaderHeight
	h.width = msg.HeaderWidth
	return h, nil
}

func (h Header) handleOptionSelectedMsg(msg OptionSelectedMsg) (Header, tea.Cmd) {
	h.pathValue = "/"
	for _, p := range msg.Path {
		h.pathValue += p.String()
	}

	return h, nil
}
