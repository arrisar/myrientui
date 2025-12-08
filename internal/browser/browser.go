package browser

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Browser struct {
	Header Header
	Footer Footer
	List   List

	loading bool
	cursor  int

	height int
	width  int
}

func New() Browser {
	b := Browser{}

	b.Header = NewHeader()
	b.Footer = NewFooter()
	b.List = NewList()

	b.loading = false
	b.cursor = 0

	b.height = 0
	b.width = 0

	return b
}

/**
 * GETTERS
 */

func (b Browser) IsLoading() bool {
	return b.loading
}

func (b Browser) Height() int {
	return b.height
}

func (b Browser) Width() int {
	return b.width
}

/**
 * SETTERS
 */

func (b *Browser) SetLoading(v bool) *Browser {
	b.loading = v
	return b
}

func (b *Browser) SetHeight(v int) *Browser {
	b.height = v
	return b
}

func (b *Browser) SetWidth(v int) *Browser {
	b.width = v
	return b
}

func (b *Browser) SetHeaderEnabled(v bool) *Browser {
	b.Header.SetEnabled(v)
	return b
}

func (b *Browser) SetListEnabled(v bool) *Browser {
	b.List.SetEnabled(v)
	return b
}

func (b *Browser) SetFooterEnabled(v bool) *Browser {
	b.Footer.SetEnabled(v)
	return b
}

/**
 * MODEL
 */

func (b Browser) Init() tea.Cmd {
	return tea.Batch(
		b.Footer.Init(),
		b.Header.Init(),
		b.List.Init(),
	)
}

func (b Browser) View() (content string) {
	content = lipgloss.JoinVertical(0, b.Header.View(), b.List.View(), b.Footer.View())
	return BrowserStyle.Height(b.height).Width(b.width).Render(content)
}

func (b Browser) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var c, cmd tea.Cmd

	b, cmd = b.handleUpdateMsg(msg)
	c = tea.Batch(c, cmd)

	b.Header, cmd = b.Header.Update(msg)
	c = tea.Batch(c, cmd)

	b.Footer, cmd = b.Footer.Update(msg)
	c = tea.Batch(c, cmd)

	b.List, cmd = b.List.Update(msg)
	c = tea.Batch(c, cmd)

	return b, c
}

/**
 * HANDLERS
 */

func (b Browser) handleUpdateMsg(msg tea.Msg) (Browser, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return b.handleKeyMsg(msg)
	case tea.WindowSizeMsg:
		return b.handleResizeMsg(msg)
	default:
		return b, nil
	}
}

func (b Browser) handleKeyMsg(msg tea.KeyMsg) (Browser, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.String() {
	case "ctrl+c":
		return b, tea.Quit
	default:
		return b, cmd
	}
}

func (b Browser) handleResizeMsg(msg tea.WindowSizeMsg) (Browser, tea.Cmd) {
	b.SetHeight(msg.Height)
	b.SetWidth(msg.Width)

	cmd := func() tea.Msg {
		BrowserHeight := b.height
		BrowserWidth := b.width

		HeaderHeight := 7
		HeaderWidth := b.width - 2
		if !b.Header.enabled {
			HeaderHeight = 0
		}

		FooterHeight := 2
		FooterWidth := b.width - 2
		if !b.Footer.enabled {
			FooterHeight = 0
		}

		ListHeight := b.height - HeaderHeight - FooterHeight - 2
		ListWidth := b.width - 2

		return BrowserResizeMsg{
			BrowserHeight,
			BrowserWidth,
			HeaderHeight,
			HeaderWidth,
			ListHeight,
			ListWidth,
			FooterHeight,
			FooterWidth,
		}
	}

	return b, cmd
}
