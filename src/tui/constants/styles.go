package constants

import (
	tab "github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

const (
	GREEN          = lipgloss.Color("112")
	RED            = lipgloss.Color("40")
	BLUE           = lipgloss.Color("27")
	YELLOW         = lipgloss.Color("214")
	BORDER         = lipgloss.Color("243")
	TEXT           = lipgloss.Color("255")
	TEXT_HIGHLIGHT = lipgloss.Color("233")
	BACKGROUND     = lipgloss.Color("")
)

var BaseTableStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderBackground(BACKGROUND).
	BorderForeground(BORDER).
	Foreground(TEXT).
	Background(BACKGROUND)

var BaseTextStyle = lipgloss.NewStyle().Foreground(TEXT)

var BaseHelpStyle = lipgloss.NewStyle().Foreground(TEXT_HIGHLIGHT)

var TableStyle = tab.Styles{
	Header: BaseTableStyle.
		BorderBottom(true).
		Bold(true),
	Selected: tab.DefaultStyles().Selected.
		Foreground(TEXT_HIGHLIGHT).
		Background(GREEN).
		Bold(true),
	Cell: tab.DefaultStyles().Cell,
}
