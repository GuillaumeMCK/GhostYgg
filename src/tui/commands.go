package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

type TickMsg time.Time

// Send a message every second.
func doTicks() tea.Cmd {
	return tea.Every(time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}
