package browser

import (
	"fmt"
	"strings"

	"github.com/arrisar/myrientui/internal/myrient"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var ContentStyle lipgloss.Style = lipgloss.NewStyle()

var MetaStyle lipgloss.Style = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#112211")).
	Padding(1, 1, 0)

type Content struct {
	screen *Screen

	options []myrient.Option
	loading bool
	list    list.Model
	path    []string
}

func (c Content) New(s *Screen) *Content {
	c.screen = s
	c.list = GetFileList()
	c.path = []string{}
	c.loading = true
	return &c
}

func (c Content) GoToCmd() tea.Cmd {
	return func() tea.Msg {
		return Scrape(fmt.Sprintf("/%s", strings.Join(c.path, "")))
	}
}

func (c Content) View() string {
	path := fmt.Sprintf("/%s", strings.Join(c.path, ""))
	body := lipgloss.NewStyle().Height(c.list.Height()).Render("")
	foot := MetaStyle.Render("(enter) open · (space) select · (a)dd tag · (c)onfiguration")

	if c.loading {
		path = MetaStyle.Render(fmt.Sprintf("%s (Loading...)", path))
	} else {
		path = MetaStyle.Render(fmt.Sprintf("%s (%d items)", path, len(c.list.Items())))
		body = c.list.View()
	}

	return lipgloss.JoinVertical(0, path, body, foot)
}

func (c Content) Init() tea.Cmd {
	return c.GoToCmd()
}

func (c Content) Update(msg tea.Msg) (*Content, tea.Cmd) {
	cmds := []tea.Cmd{}
	navigating := false
	switch msg := msg.(type) {

	// scrape results
	case ScrapeResultMsg:
		c.loading = false
		c.list.SetItems([]list.Item{})

		offset := 0
		if len(c.path) > 0 {
			offset = 1
			c.list.InsertItem(0, FileItem("../"))
		}

		for i, option := range msg.options {
			c.list.InsertItem(i+offset, FileItem(option.Label))
		}

	// key input
	case tea.KeyMsg:
		if c.loading == true {
			break
		}

		switch msg.String() {
		case "enter":
			item := c.list.SelectedItem().(FileItem)
			part := fmt.Sprint(item)

			switch true {
			case part == "../":
				c.path = c.path[:len(c.path)-1]
				navigating = true

			case strings.HasSuffix(part, ".zip"):
				navigating = false

			default:
				c.path = append(c.path, part)
				navigating = true
			}
		}

	// window resize
	case tea.WindowSizeMsg:
		c.list.SetHeight(msg.Height - 8)
	}

	if navigating {
		c.loading = true
		cmds = append(cmds, c.GoToCmd())
	}

	list, cmd := c.list.Update(msg)
	c.list = list

	cmds = append(cmds, cmd)
	return &c, tea.Batch(cmds...)
}
