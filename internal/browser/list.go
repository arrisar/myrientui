package browser

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sahilm/fuzzy"
)

type List struct {
	enabled bool
	loading bool

	cursor  int
	options Options
	path    Options

	filterActive bool
	filterValue  string

	height int
	width  int
}

func NewList() List {
	l := List{}

	l.enabled = true
	l.loading = true

	l.cursor = 0
	l.options = Options{}
	l.path = Options{}

	l.filterActive = false
	l.filterValue = ""

	l.height = 0
	l.width = 0

	return l
}

/**
 * GETTERS
 */

func (l List) IsEnabled() bool {
	return l.enabled
}

func (l List) FilteredOptions() Options {
	if len(l.filterValue) == 0 {
		return l.options
	}

	f := Options{}
	if len(l.path) > 0 {
		f = Options{"../"}
	}

	results := fuzzy.FindFrom(l.filterValue, l.options)
	for _, r := range results {
		f = append(f, Option(r.Str))
	}

	return f
}

/**
 * SETTERS
 */

func (l *List) SetEnabled(v bool) *List {
	l.enabled = v
	return l
}

/**
 * MODEL
 */

func (l List) Init() tea.Cmd {
	return func() tea.Msg {
		return OptionSelectedMsg{
			Option: Option(""),
			Path:   l.path,
		}
	}
}

func (l List) View() string {
	content := ""
	filtered := l.FilteredOptions()

	pageSize := l.height - 3
	for i, o := range filtered {
		if i >= pageSize {
			break
		}

		prefix := "  ./"
		itemStyle := ListItemStyle

		if i == l.cursor {
			prefix = "> ./"
			itemStyle = ListCurrentStyle
		}

		if i > 0 {
			prefix = "\n" + prefix
		}

		content = content + itemStyle.Render(prefix+string(o))
	}

	return ListStyle.Height(l.height).Width(l.width).Render(content)
}

func (l List) Update(msg tea.Msg) (List, tea.Cmd) {
	var c, cmd tea.Cmd

	l, cmd = l.handleUpdateMsg(msg)
	c = tea.Batch(c, cmd)

	return l, tea.Batch(c)
}

/**
 * HANDLERS
 */

func (l List) handleUpdateMsg(msg tea.Msg) (List, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return l.handleKeyMsg(msg)
	case BrowserResizeMsg:
		return l.handleBrowserResizeMsg(msg)
	case FilterStateMsg:
		return l.handleFilterStateMsg(msg)
	case OptionsChangedMsg:
		return l.handleOptionsChangedMsg(msg)
	case OptionsLoadingMsg:
		return l.handleOptionsLoadingMsg(msg)
	default:
		return l, nil
	}
}

func (l List) handleKeyMsg(msg tea.KeyMsg) (List, tea.Cmd) {
	if l.filterActive || l.loading {
		return l, nil
	}

	var cmd tea.Cmd
	switch msg.String() {

	case "enter":
		o := l.FilteredOptions()[l.cursor]
		p := l.path

		back := o.String() == "../"
		depth := len(p)
		switch true {
		case back && depth == 0:
			return l, nil
		case back && depth == 1:
			o = ""
			p = Options{}
		case back && depth > 1:
			prev := p[depth-1]
			o = prev
			p = p[0 : depth-1]
		default:
			p = append(p, o)
		}

		if o.IsFile() {
			fmt.Println("FILE CHOSEN")
			return l, nil
		}

		l.path = p
		cmd = func() tea.Msg {
			return OptionSelectedMsg{
				Option: o,
				Path:   p,
			}
		}

	case "up":
		if l.cursor > 0 {
			l.cursor -= 1
		}

	case "down":
		if l.cursor < len(l.options)-1 {
			l.cursor += 1
		}
	}

	return l, cmd
}

func (l List) handleBrowserResizeMsg(msg BrowserResizeMsg) (List, tea.Cmd) {
	l.height = msg.ListHeight
	l.width = msg.ListWidth
	return l, nil
}

func (l List) handleFilterStateMsg(msg FilterStateMsg) (List, tea.Cmd) {
	l.filterActive = msg.Active
	l.filterValue = msg.Value

	if msg.Active {
		l.cursor = -1
	} else {
		l.cursor = 0
	}

	return l, nil
}

func (l List) handleOptionsChangedMsg(msg OptionsChangedMsg) (List, tea.Cmd) {
	l.loading = false
	l.options = msg.Options

	if len(l.path) > 0 {
		l.options = append(Options{"../"}, l.options...)
	}

	if l.cursor != 0 {
		l.cursor = 0
	}

	return l, nil
}

func (l List) handleOptionsLoadingMsg(OptionsLoadingMsg) (List, tea.Cmd) {
	l.loading = true
	return l, nil
}
