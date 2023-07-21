package tui

import (
	"GhostYgg/src/tui/constants"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type Help struct {
	keys      constants.KeyMap
	help      help.Model
	maxHeight int
}

// NewHelp creates a new help model.
func NewHelp() Help {
	h := help.New()
	h.Styles = constants.HelpStyle

	return Help{
		keys:      constants.Keys,
		help:      h,
		maxHeight: len(constants.Keys.FullHelp()[0]),
	}
}

func (m Help) Init() tea.Cmd { return nil }

func (m Help) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Help) View() string {
	return constants.BaseHelpStyle.Render(m.help.View(m.keys))
}

func (m *Help) switchHelp() {
	m.help.ShowAll = !m.help.ShowAll
	constants.HelpHeight = m.GetHeight()
}

func (m Help) GetHeight() int {
	if !m.help.ShowAll {
		return 1
	}
	return m.maxHeight
}
