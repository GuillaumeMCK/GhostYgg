package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

type TickMsg time.Time

func doTicks() tea.Cmd {
	return tea.Every(time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

type UpdateTableMsg struct{}

func forceUpdateTable() tea.Cmd {
	return func() tea.Msg {
		return UpdateTableMsg{}
	}
}

type SelectedRowMsg struct {
	Index int
}

func selectedRow(index int) tea.Cmd {
	return func() tea.Msg {
		return SelectedRowMsg{Index: index}
	}
}

type UpdateHeightMsg struct {
	Height int
}
