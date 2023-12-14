package tui

import (
	"github.com/GuillaumeMCK/GhostYgg/src/tui/constants"
	"github.com/GuillaumeMCK/GhostYgg/src/utils"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Help represents the TUI help model.
type Help struct {
	keys      constants.KeyMap
	help      *help.Model
	maxHeight int
	size      *utils.Size
}

func (h *Help) Init() tea.Cmd { return nil }

func (h *Help) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.Keys.Help):
			h.Swicth()
			cmd = updateContainer()
		}
	}
	return h, cmd
}

func (h *Help) View() string {
	h.help.Width = h.size.Width - 4
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(constants.BORDER).
		Width(h.size.Width-2).
		Align(lipgloss.Center).
		Render(constants.BaseHelpStyle.Render(h.help.View(h.keys))) + "\n"
}

func (h *Help) Swicth() {
	h.help.ShowAll = !h.help.ShowAll
}

func (h *Help) getHeight() int {
	if !h.help.ShowAll {
		return 3
	}
	return h.maxHeight + 2
}

// NewHelp creates a new help model.
func NewHelp(size *utils.Size) *Help {
	h := help.New()
	h.Styles = constants.HelpStyle

	return &Help{
		keys:      constants.Keys,
		help:      &h,
		maxHeight: len(constants.Keys.FullHelp()[0]),
		size:      size,
	}
}
