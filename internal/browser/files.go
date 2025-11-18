package browser

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func GetFileList() list.Model {
	model := list.New([]list.Item{}, FileDelegate{}, 0, 0)
	model.SetFilteringEnabled(true)
	model.SetShowFilter(false)
	model.SetShowHelp(false)
	model.SetShowPagination(false)
	model.SetShowStatusBar(false)
	model.SetShowTitle(false)
	return model
}

var FileItemStyle lipgloss.Style = lipgloss.NewStyle()

var FileSelectedStyle lipgloss.Style = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#117711"))

/**
 * FileItem
 */

type FileItem string

func (f FileItem) FilterValue() string {
	return string(f)
}

/**
 * FileDelegate
 */

type FileDelegate struct{}

func (f FileDelegate) Height() int {
	return 1
}

func (f FileDelegate) Spacing() int {
	return 0
}

func (f FileDelegate) Update(tea.Msg, *list.Model) tea.Cmd {
	return nil
}

func (f FileDelegate) Render(writer io.Writer, list list.Model, index int, item list.Item) {
	i, ok := item.(FileItem)
	if !ok {
		return
	}

	str := fmt.Sprintf("./%s", i)
	fn := func(s ...string) string {
		return FileItemStyle.Render("  " + strings.Join(s, " "))
	}

	if index == list.Index() {
		fn = func(s ...string) string {
			return FileSelectedStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(writer, fn(str))
}
