package tui

import (
	"GhostYgg/src/tui/constants"
	hlp "github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type Help struct {
	keys      constants.KeyMap
	help      hlp.Model
	maxHeight int
}

// NewHelp creates a new help model.
func NewHelp() Help {
	return Help{
		keys:      constants.Keys,
		help:      hlp.New(),
		maxHeight: len(constants.Keys.FullHelp()[0]) + 1,
	}
}

func (m Help) Init() tea.Cmd { return nil }

func (m Help) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "?", "h":
			m.help.ShowAll = !m.help.ShowAll
		}
	}
	return m, nil
}

func (m Help) View() string {
	return constants.BaseHelpStyle.Render(m.help.View(m.keys))
}

func (m Help) GetHeight() int {
	if m.help.ShowAll {
		return 2
	}
	return m.maxHeight + 1
}
