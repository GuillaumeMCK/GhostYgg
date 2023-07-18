package constants

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

const (
	GREEN      = lipgloss.Color("112")
	RED        = lipgloss.Color("196")
	BLUE       = lipgloss.Color("27")
	YELLOW     = lipgloss.Color("214")
	BORDER     = lipgloss.Color("240")
	TEXT       = lipgloss.Color("255")
	DESC       = lipgloss.Color("245")
	HIGHLIGHT  = lipgloss.Color("232")
	BACKGROUND = lipgloss.Color("")
)

var HelpStyle = help.Styles{
	FullKey:        lipgloss.NewStyle().Foreground(TEXT),
	ShortKey:       lipgloss.NewStyle().Foreground(TEXT),
	FullDesc:       lipgloss.NewStyle().Foreground(DESC),
	ShortDesc:      lipgloss.NewStyle().Foreground(DESC),
	FullSeparator:  lipgloss.NewStyle().Foreground(GREEN),
	ShortSeparator: lipgloss.NewStyle().Foreground(GREEN),
}

var BaseHelpStyle = lipgloss.NewStyle().
	BorderBackground(BACKGROUND).
	Padding(0, 1)

var BaseTableStyle = lipgloss.NewStyle().
	BorderBackground(BACKGROUND).
	Background(BACKGROUND).
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(BORDER).
	Foreground(TEXT)

var TableStyle = table.Styles{
	Header: table.DefaultStyles().Header.
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(BORDER).
		BorderBackground(BACKGROUND).
		Background(BACKGROUND).
		BorderBottom(true).
		Foreground(TEXT).
		Bold(true),
	Selected: table.DefaultStyles().Selected.
		Foreground(HIGHLIGHT).
		Background(GREEN).
		Bold(true),
	Cell: table.DefaultStyles().Cell,
}
