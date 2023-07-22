package tui

import (
	"GhostYgg/src/client"
	"GhostYgg/src/tui/constants"
	"GhostYgg/src/utils"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// Model represents the TUI model.
type Model struct {
	table         Table
	help          Help
	torrentClient client.Model
}

// Init initializes the TUI model and returns a command to execute during the initialization.
func (m Model) Init() tea.Cmd {
	return tea.Batch(tea.ClearScreen, tea.ClearScrollArea, updateTuiLoop())
}

// Update handles incoming messages and updates the TUI model accordingly.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case UpdateTuiLoopMsg, tea.WindowSizeMsg:
		if winSize, ok := msg.(tea.WindowSizeMsg); ok {
			constants.WindowSize = winSize
			cmd = tea.ClearScreen
		}
		return m, tea.Batch(cmd, updateTable(), updateTuiLoop())
	case UpdateTableMsg:
		m.table.refresh(*m.torrentClient.DownloadsQueue)
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
			(*m.torrentClient.DownloadsQueue)[m.table.selectedRow()].Abort()
			return m, nil
		case key.Matches(msg, constants.Keys.PauseAndPlay):
			(*m.torrentClient.DownloadsQueue)[m.table.selectedRow()].PauseAndPlay()
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
func (m Model) View() string {
	return m.table.View() + "\n" + m.help.View()
}

// New creates a new TUI model and returns it along with a command to execute during initialization.
func New(torrentFiles []string) (tea.Model, tea.Cmd) {
	torrentClient := client.New(constants.DownloadFolder, torrentFiles)
	err := torrentClient.Start()
	if err != nil {
		panic("error starting torrent client")
	}

	m := Model{
		table: NewTable(&TableCtx{
			Rows:    make([]client.DownloadInfos, 0),
			Columns: constants.TableColumns,
			Widths:  constants.TableWidths,
		}),
		help:          NewHelp(),
		torrentClient: torrentClient,
	}

	return m, nil
}
