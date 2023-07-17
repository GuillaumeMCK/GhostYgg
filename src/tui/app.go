package tui

import (
	"GhostYgg/src/client"
	"GhostYgg/src/tui/constants"
	"GhostYgg/src/utils"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	table         Table
	help          Help
	selectedRow   int
	torrentClient *client.Model
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(doTicks())
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		constants.WindowSize = msg
	case TickMsg:
		// Pass the tick message to the table. This will trigger a table update
		m.Update(UpdateTableMsg{})
		return m, nil
	case SelectedRowMsg:
		m.selectedRow = msg.Index
		return m, nil
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.Keys.Help):
			m.help.Update(msg)
			return m, nil
		//case key.Matches(msg, constants.Keys.Add):
		//	// TODO: add a new download. use sqweek/dialog lib to pick a file
		case key.Matches(msg, constants.Keys.Open):
			utils.OpenDirectory(constants.DownloadFolder)
			return m, nil
		case key.Matches(msg, constants.Keys.Delete):
			m.torrentClient.DownloadsInfos[m.selectedRow].Abort()
			return m.Update(UpdateTableMsg{})
		case key.Matches(msg, constants.Keys.PauseAndPlay):
			m.torrentClient.DownloadsInfos[m.selectedRow].PauseAndPlay()
			return m.Update(msg)
		case key.Matches(msg, constants.Keys.Quit):
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Model) View() string {
	return m.View() + m.help.View()
}

func App() (tea.Model, tea.Cmd) {
	m := Model{
		table: NewTable(&TableCtx{
			Rows:    constants.TableRows,
			Columns: constants.TableColumns,
			Widths:  constants.TableWidths,
		}),
		help:          NewHelp(),
		selectedRow:   0,
		torrentClient: client.New(constants.DownloadFolder, constants.TorrentFiles),
	}
	return m, nil
}
