package help

import (
	"GhostYgg/src/tui/constants"
	hlp "github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	keys      constants.KeyMap
	help      hlp.Model
	maxHeight int
}

func New() Model {
	return Model{
		keys:      constants.Keys,
		help:      hlp.New(),
		maxHeight: len(constants.Keys.FullHelp()[0]) + 1,
	}
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m Model) View() string {
	return constants.BaseHelpStyle.Render(m.help.View(m.keys))
}

func (m Model) GetHeight() int {
	if m.help.ShowAll {
		return 2
	}
	return m.maxHeight + 1
}
