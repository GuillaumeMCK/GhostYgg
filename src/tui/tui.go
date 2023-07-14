package tui

import (
	"GhostYgg/src/tui/constants"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

// StartTea the entry point for the UI. Initializes the model.
func StartTea(torrentFiles []string, downloadFolder string) error {
	constants.DownloadFolder = downloadFolder
	constants.TorrentFiles = torrentFiles
	m, _ := App()
	constants.P = tea.NewProgram(m, tea.WithAltScreen())
	if _, err := constants.P.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
	return nil
}
