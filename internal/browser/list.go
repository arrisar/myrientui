package browser

import (
	"fmt"
	"math"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sahilm/fuzzy"
)

type List struct {
	enabled bool
	loading bool
	spinner spinner.Model

	cursor   int
	options  Options
	filtered Options
	path     Options

	filterEnabled bool
	filterActive  bool
	filterApplied bool
	filterValue   string

	height int
	width  int
}

func NewList() List {
	l := List{}

	l.enabled = true
	l.loading = true
	l.spinner = spinner.New()

	l.cursor = 0
	l.options = Options{}
	l.path = Options{}

	l.filterEnabled = true
	l.filterActive = false
	l.filterApplied = false
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
	f := Options{}
	if len(l.path) > 0 {
		f = Options{"../"}
	}

	if len(l.filterValue) == 0 {
		return append(f, l.options...)
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
	return tea.Batch(
		l.spinner.Tick,
		func() tea.Msg {
			return OptionSelectedMsg{
				Option: Option(""),
				Path:   l.path,
			}
		})
}

func (l List) View() string {
	pageSize := l.height - 4
	pageMiddle := int(math.Ceil(float64(pageSize / 2)))
	lastIndex := len(l.filtered) - 1

	// determine offset
	offset := max(min(l.cursor-pageMiddle, lastIndex-pageSize), 0)
	pageLastIndex := min(offset+pageSize, lastIndex, len(l.filtered)-1)

	files := ""
	for i, o := range l.filtered {
		if i < offset {
			continue
		}

		if i > pageLastIndex {
			break
		}

		prefix := "  ./"
		itemStyle := ListItemStyle

		if i == l.cursor {
			prefix = "> ./"
			itemStyle = ListCurrentStyle
		}

		if i > offset {
			prefix = "\n" + prefix
		}

		files = files + itemStyle.Render(fmt.Sprint(prefix, o))
	}

	// counts row
	count := fmt.Sprintf("%d/%d items", l.cursor+1, len(l.filtered))
	if l.loading {
		count = "Loading items " + l.spinner.View()
	}

	// render
	files = ListStyle.Height(l.height - 2).Width(l.width).Render(files)
	count = ListCountStyle.Render(count)
	return lipgloss.JoinVertical(0, files, count)
}

func (l List) Update(msg tea.Msg) (List, tea.Cmd) {
	var c, cmd tea.Cmd

	l, cmd = l.handleUpdateMsg(msg)
	c = tea.Batch(c, cmd)

	l.spinner, cmd = l.spinner.Update(msg)
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

	case "enter", "right":
		l, cmd = l.handleSelectOption(l.filtered[l.cursor])

	case "backspace", "left":
		l, cmd = l.handleSelectOption("../")

	case "up":
		if l.cursor > 0 {
			l.cursor -= 1
		}

	case "down":
		if l.cursor < len(l.filtered)-1 {
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
	l.filterEnabled = msg.Enabled
	l.filterActive = msg.Active
	l.filterApplied = msg.Applied
	l.filterValue = msg.Value
	l.filtered = l.FilteredOptions()

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
	l.filtered = l.FilteredOptions()

	if l.cursor != 0 {
		l.cursor = 0
	}

	return l, nil
}

func (l List) handleOptionsLoadingMsg(OptionsLoadingMsg) (List, tea.Cmd) {
	l.loading = true
	l.filtered = Options{}
	l.options = Options{}
	return l, nil
}

func (l List) handleSelectOption(t Option) (List, tea.Cmd) {
	o := t
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
		return l, nil
	}

	l.path = p
	cmd := tea.Batch(
		func() tea.Msg {
			return FilterStateMsg{
				Enabled: l.filterEnabled,
				Active:  false,
				Applied: false,
				Value:   "",
			}
		},
		func() tea.Msg {
			return OptionSelectedMsg{
				Option: o,
				Path:   p,
			}
		},
	)

	return l, cmd
}
