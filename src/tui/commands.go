package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

type UpdateTuiLoopMsg struct{}

func updateTuiLoop() tea.Cmd {
	return tea.Every(time.Millisecond*150, func(time.Time) tea.Msg {
		return UpdateTuiLoopMsg{}
	})
}

type UpdateTableMsg struct{}

func updateTable() tea.Cmd {
	return func() tea.Msg {
		return UpdateTableMsg{}
	}
}
