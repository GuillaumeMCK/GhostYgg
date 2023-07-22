package tui

import (
	"GhostYgg/src/tui/constants"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

// Help represents the TUI help model.
type Help struct {
	keys      constants.KeyMap
	help      help.Model
	maxHeight int
}

// NewHelp creates a new help model.
func NewHelp() *Help {
	h := help.New()
	h.Styles = constants.HelpStyle

	return &Help{
		keys:      constants.Keys,
		help:      h,
		maxHeight: len(constants.Keys.FullHelp()[0]),
	}
}

func (h *Help) Init() tea.Cmd { return nil }

func (h *Help) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return h, nil
}

func (h *Help) View() string {
	return constants.BaseHelpStyle.Render(h.help.View(h.keys))
}

func (h *Help) switchHelp() {
	h.help.ShowAll = !h.help.ShowAll
	constants.HelpHeight = h.getHeight()
}

func (h *Help) getHeight() int {
	if !h.help.ShowAll {
		return 1 + 1
	}
	return h.maxHeight + 1
}
