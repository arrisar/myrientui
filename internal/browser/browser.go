package browser

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Browser struct {
	screen *Screen
	ready  bool
}

func Start() {
	b := Browser{}
	b.screen = Screen{}.New(&b)
	p := tea.NewProgram(b, tea.WithAltScreen(), tea.WithMouseCellMotion())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func (b Browser) Init() tea.Cmd {
	return b.screen.Init()
}

func (b Browser) View() string {
	return b.screen.View()
}

func (b Browser) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// key presses
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return b, tea.Quit
		}

	// window resize
	case tea.WindowSizeMsg:
		b.screen.height = msg.Height
		b.screen.width = msg.Width
		b.ready = true
	}

	s, cmd := b.screen.Update(msg)
	b.screen = s
	return b, cmd
}
