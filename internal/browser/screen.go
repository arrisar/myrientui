package browser

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var ScreenStyle lipgloss.Style = lipgloss.NewStyle()

type Screen struct {
	browser *Browser
	header  *Header
	content *Content

	height int
	width  int
}

func (s Screen) New(b *Browser) *Screen {
	s.browser = b
	s.header = Header{}.New(&s)
	s.content = Content{}.New(&s)
	return &s
}

func (s Screen) View() string {
	return ScreenStyle.Render(lipgloss.JoinVertical(0, s.header.View(), s.content.View()))
}

func (s Screen) Init() tea.Cmd {
	return tea.Batch(s.header.Init(), s.content.Init())
}

func (s Screen) Update(msg tea.Msg) (*Screen, tea.Cmd) {
	hm, hc := s.header.Update(msg)
	cm, cc := s.content.Update(msg)

	s.header = hm
	s.content = cm

	return &s, tea.Batch(hc, cc)
}
