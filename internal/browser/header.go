package browser

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var HeaderStyle lipgloss.Style = func() lipgloss.Style {
	b := lipgloss.RoundedBorder()
	b.Right = "├"
	return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
}()

type Header struct {
	screen *Screen
	title  string
}

func (h Header) New(s *Screen) *Header {
	h.screen = s
	h.title = "MyrienTUI"
	return &h
}

func (h Header) View() string {
	title := HeaderStyle.Render(h.title)
	line := strings.Repeat("─", max(0, h.screen.width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (h Header) Init() tea.Cmd {
	return nil
}

func (h Header) Update(msg tea.Msg) (*Header, tea.Cmd) {
	return &h, nil
}
