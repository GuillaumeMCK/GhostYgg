package tui

import (
	"GhostYgg/src/client"
	"GhostYgg/src/tui/constants"
	"GhostYgg/src/utils"
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// TUI represents the TUI model
type TUI struct {
	table         *Table
	help          *Help
	torrentClient *client.Model
}

// Init initializes the TUI model and returns a command to execute during the initialization.
func (m TUI) Init() tea.Cmd {
	return tea.Batch(tea.ClearScreen, tea.ClearScrollArea, updateTuiLoop())
}

// Update handles incoming messages and updates the TUI model accordingly.
func (m TUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case UpdateTuiLoopMsg, tea.WindowSizeMsg:
		if winSize, ok := msg.(tea.WindowSizeMsg); ok {
			constants.WindowSize = winSize
			cmd = tea.ClearScreen
		}
		return m, tea.Batch(cmd, updateTable(), updateTuiLoop())
	case UpdateTableMsg:
		m.table.refresh(*m.torrentClient.Torrents)
		return m, nil
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.Keys.Help):
			m.help.switchHelp()
			return m, updateTable()
		//case key.Matches(msg, constants.Keys.Add):
		//	// TODO: add a new download. use sqweek/dialog.go lib to pick a file
		case key.Matches(msg, constants.Keys.Open):
			utils.OpenDirectory(constants.DownloadFolder)
			return m, nil
		case key.Matches(msg, constants.Keys.Delete):
			(*m.torrentClient.Torrents)[m.table.selectedRow()].Abort()
			return m, nil
		case key.Matches(msg, constants.Keys.PauseAndPlay):
			(*m.torrentClient.Torrents)[m.table.selectedRow()].PauseAndPlay()
			return m, nil
		case key.Matches(msg, constants.Keys.Quit):
			m.torrentClient.Abort()
			return m, tea.Quit
		}
	}
	m.table.table, cmd = m.table.table.Update(msg)
	return m, cmd
}

// View renders the TUI view as a string.
func (m TUI) View() string {
	return m.table.View() + "\n" + m.help.View()
}

// NewTUI creates a new TUI model.
func NewTUI(torrentFiles []string) (tea.Model, tea.Cmd) {
	torrentClient, _ := client.New(constants.DownloadFolder, torrentFiles)

	err := torrentClient.Start()
	if err != nil {
		panic("error starting torrent client")
	}

	m := TUI{
		table: NewTable(&TableCtx{
			Rows:    make([]client.TorrentInfos, 0),
			Columns: constants.TableColumns,
			Widths:  constants.TableWidths,
		}),
		help:          NewHelp(),
		torrentClient: torrentClient,
	}

	return m, nil
}

// StartTUI starts the TUI.
func StartTUI(torrentFiles []string, downloadFolder string) error {
	constants.DownloadFolder = downloadFolder
	m, err := NewTUI(torrentFiles)
	if err != nil {
		return fmt.Errorf("error creating TUI: %v", err)
	}

	constants.P = tea.NewProgram(m, tea.WithAltScreen())
	if _, err := constants.P.Run(); err != nil {
		return fmt.Errorf("error running program: %v", err)
	}

	return nil
}
