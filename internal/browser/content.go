package browser

import (
	"fmt"
	"strings"

	"github.com/arrisar/myrientui/internal/myrient"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var ContentStyle lipgloss.Style = lipgloss.NewStyle()

var HeadStyle lipgloss.Style = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#112211")).
	Padding(1)

var FootStyle lipgloss.Style = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#112211")).
	Padding(1, 1, 0)

type Content struct {
	screen *Screen

	options   []myrient.Option
	loading   bool
	filtering bool
	filter    textinput.Model
	list      list.Model
	path      []string
}

func (c Content) New(s *Screen) *Content {
	c.screen = s
	c.filter = GetFilterInput()
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
	path := fmt.Sprintf("Path: /%s", strings.Join(c.path, ""))
	head := HeadStyle.Render(path)
	body := lipgloss.NewStyle().Height(c.list.Height()).Render("")
	foot := "(enter) directory · (f)ilter · (a)dd tag · (c)onfiguration"

	switch true {
	case c.loading:
		foot = "Loading..."
	case c.filtering:
		head = c.filter.View() + "\n"
		body = c.list.View()
		foot = "(enter) continue · (esc) clear"
	case len(c.filter.Value()) > 0:
		head = c.filter.View() + "\n"
		body = c.list.View()
	default:
		head = HeadStyle.Render(fmt.Sprintf("%s (%d items)", path, len(c.list.Items())))
		body = c.list.View()
	}

	foot = FootStyle.Render(foot)
	return lipgloss.JoinVertical(0, head, body, foot)
}

func (c Content) Init() tea.Cmd {
	return c.GoToCmd()
}

func (c Content) Update(msg tea.Msg) (*Content, tea.Cmd) {
	cmds := []tea.Cmd{}
	navigating := false
	forwardToList := false
	forwardToFilter := false

	switch msg := msg.(type) {

	// scrape results
	case ScrapeResultMsg:
		forwardToList = true
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
		switch true {
		case c.loading:
			break
		case c.filtering:
			applied := len(c.filter.Value()) > 0
			switch msg.String() {
			case "enter":
				c.filtering = false
				c.filter.Blur()

				if applied {
					c.list.SetFilterState(list.FilterApplied)
				} else {
					c.list.SetFilterState(list.Unfiltered)
				}
			case "esc":
				forwardToFilter = true
				c.filtering = false
				c.filter.Blur()
				c.filter.SetValue("")
				c.list.SetFilterState(list.Unfiltered)
			default:
				forwardToFilter = true
				if applied {
					c.list.SetFilterState(list.Filtering)
				} else {
					c.list.SetFilterState(list.Unfiltered)
				}
			}
		default:
			switch msg.String() {
			case "up", "down":
				forwardToList = true
			case "enter":
				item := c.list.SelectedItem().(FileItem)
				part := fmt.Sprint(item)

				switch true {
				case part == "../":
					c.path = c.path[:len(c.path)-1]
					navigating = true

				case strings.HasSuffix(part, ".zip"):
					navigating = false
					forwardToList = true

				default:
					c.path = append(c.path, part)
					navigating = true
				}
			case "f":
				c.filtering = true
				c.filter.Focus()
			case "c":
				break
			case "a":
				break
			}
		}

	// window resize
	case tea.WindowSizeMsg:
		c.list.SetHeight(msg.Height - 9)
		forwardToList = true
	}

	// go to new page
	if navigating {
		c.filter.Reset()
		c.list.ResetFilter()
		c.loading = true
		cmds = append(cmds, c.GoToCmd())
	}

	// forward msg to filter input
	if forwardToFilter {
		filter, cmd := c.filter.Update(msg)
		c.filter = filter
		c.list.SetFilterText(c.filter.Value())
		cmds = append(cmds, cmd)
	}

	// forward msg to list
	if forwardToList {
		list, cmd := c.list.Update(msg)
		c.list = list
		cmds = append(cmds, cmd)
	}

	return &c, tea.Batch(cmds...)
}
