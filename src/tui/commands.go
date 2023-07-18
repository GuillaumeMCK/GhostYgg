package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

type UpdateTuiMsg time.Time

func updateTui() tea.Cmd {
	return tea.Tick(time.Millisecond*150, func(time.Time) tea.Msg {
		return UpdateTuiMsg{}
	})
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
