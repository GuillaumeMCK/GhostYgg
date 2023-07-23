package constants

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

var (
	GREEN      = lipgloss.AdaptiveColor{Light: "112", Dark: "112"}
	RED        = lipgloss.AdaptiveColor{Light: "196", Dark: "196"}
	BLUE       = lipgloss.AdaptiveColor{Light: "27", Dark: "27"}
	YELLOW     = lipgloss.AdaptiveColor{Light: "214", Dark: "214"}
	BORDER     = lipgloss.AdaptiveColor{Light: "0", Dark: "240"}
	TEXT       = lipgloss.AdaptiveColor{Light: "0", Dark: "255"}
	DESC       = lipgloss.AdaptiveColor{Light: "244", Dark: "244"}
	HIGHLIGHT  = lipgloss.AdaptiveColor{Light: "232", Dark: "232"}
	BACKGROUND = lipgloss.AdaptiveColor{Light: "", Dark: ""}
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
