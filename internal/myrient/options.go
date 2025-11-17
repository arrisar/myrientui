package myrient

import (
	gloss "github.com/charmbracelet/lipgloss"
)

var OptionStyle gloss.Style = gloss.NewStyle().
	Foreground(gloss.Color("red"))

type Option struct {
	IsDir bool
	Date  string
	Link  string
	Label string
	Size  string
}
