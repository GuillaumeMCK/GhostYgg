package tui

import (
	constants2 "github.com/GuillaumeMCK/GhostYgg/tui/constants"
	"github.com/GuillaumeMCK/GhostYgg/utils"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Help represents the TUI help model.
type Help struct {
	keys      constants2.KeyMap
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
		case key.Matches(msg, constants2.Keys.Help):
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
		BorderForeground(constants2.BORDER).
		Width(h.size.Width-2).
		Align(lipgloss.Center).
		Render(constants2.BaseHelpStyle.Render(h.help.View(h.keys))) + "\n"
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
	h.Styles = constants2.HelpStyle

	return &Help{
		keys:      constants2.Keys,
		help:      &h,
		maxHeight: len(constants2.Keys.FullHelp()[0]),
		size:      size,
	}
}
