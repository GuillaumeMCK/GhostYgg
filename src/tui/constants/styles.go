package constants

import (
	"github.com/charmbracelet/bubbles/filepicker"
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
	HIGHLIGHT  = lipgloss.AdaptiveColor{Light: "252", Dark: "232"}
	BACKGROUND = lipgloss.AdaptiveColor{Light: "", Dark: ""}
	HEADER     = [3]string{"#87D700", "#11D700", "#00D795"}
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
	BorderForeground(BORDER).
	Foreground(TEXT)

var BaseContainerStyle = BaseTableStyle.Copy().BorderStyle(lipgloss.RoundedBorder())

var TableStyle = table.Styles{
	Header: table.DefaultStyles().Header.
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(BORDER).
		BorderBackground(BACKGROUND).
		Background(BACKGROUND).
		//BorderBottom(true).
		BorderTop(true).
		Foreground(TEXT).
		Bold(true),
	Selected: table.DefaultStyles().Selected.
		Foreground(GREEN).
		//Background(GREEN).
		Bold(true),
	Cell: table.DefaultStyles().Cell,
}

var FilePickerStyle = filepicker.Styles{
	Selected: filepicker.DefaultStyles().Selected.
		Foreground(GREEN).
		Bold(true),
	DisabledSelected: filepicker.DefaultStyles().DisabledSelected.
		Foreground(RED).
		Bold(true),
	DisabledCursor: filepicker.DefaultStyles().DisabledCursor.
		Foreground(RED).
		Bold(true),
	Cursor: filepicker.DefaultStyles().Cursor.
		Foreground(GREEN).
		Bold(true),
	File: filepicker.DefaultStyles().File.
		Foreground(TEXT).
		Bold(true),
	DisabledFile: filepicker.DefaultStyles().DisabledFile.
		Foreground(DESC).
		Bold(true),
	Symlink: filepicker.DefaultStyles().Symlink.
		Foreground(BLUE).
		Bold(true),
	Permission: filepicker.DefaultStyles().Permission.
		Foreground(DESC).
		Bold(false),
	FileSize: filepicker.DefaultStyles().FileSize.
		Foreground(DESC).
		Bold(true),
	EmptyDirectory: filepicker.DefaultStyles().EmptyDirectory.
		Foreground(DESC).
		Bold(false),
}

var HeadStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(HIGHLIGHT).
	Background(TEXT)
