package tui

import (
	"github.com/GuillaumeMCK/GhostYgg/tui/constants"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

type Header struct{}

func (m Header) Init() tea.Cmd {
	return nil
}

func (m Header) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Header) View() string {
	var s strings.Builder

	s.WriteString(constants.HeadStyle.Render("Ghost"))

	for i, char := range "Ygg" {
		s.WriteString(lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#000000")).
			Background(lipgloss.Color(constants.HEADER[i%3])).
			Italic(true).
			Render(string(char)))
	}

	// Write ghost emoji to the end of the header.
	s.WriteString(lipgloss.NewStyle().
		Background(lipgloss.Color("")).
		Render(" 👻"))

	centeredHeader := lipgloss.NewStyle().
		Align(lipgloss.Left).
		PaddingLeft(1).
		Render(s.String())

	return centeredHeader + "\n"
}

func (m Header) getHeight() int {
	return 2
}

func NewHeader() *Header {
	return &Header{}
}
