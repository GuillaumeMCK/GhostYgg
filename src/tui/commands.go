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

type UpdateContainerMsg struct{}

func updateContainer() tea.Cmd {
	return func() tea.Msg {
		return UpdateContainerMsg{}
	}
}

type ClearErrorMsg struct{}

func clearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return ClearErrorMsg{}
	})
}

type AddTorrentMsg struct {
	Path string
}

func addTorrent(path string) tea.Cmd {
	return func() tea.Msg {
		return AddTorrentMsg{Path: path}
	}
}
