package table

import tea "github.com/charmbracelet/bubbletea"

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
