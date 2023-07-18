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
	tableCtx      TableCtx
	help          Help
	selectedRow   int
	torrentClient client.Model
}

func (m Model) Init() tea.Cmd {
	return updateTui()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case UpdateTuiMsg, tea.WindowSizeMsg:
		cmds := []tea.Cmd{updateTui()}
		if msg, ok := msg.(tea.WindowSizeMsg); ok {
			constants.WindowSize = msg
			cmds = append(cmds, tea.ClearScreen)
		}
		m.tableCtx.Rows = *m.torrentClient.DownloadsQueue
		m.table.Update(cmds)
		return m, tea.Batch(cmds...)
	case SelectedRowMsg:
		m.selectedRow = msg.Index
		return m, nil
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.Keys.Help):
			m.help.Update(msg)
			return m, nil
		//case key.Matches(msg, constants.Keys.Add):
		//	// TODO: add a new download. use sqweek/dialog.go lib to pick a file
		case key.Matches(msg, constants.Keys.Open):
			utils.OpenDirectory(constants.DownloadFolder)
			return m, nil
		case key.Matches(msg, constants.Keys.Delete):
			(*m.torrentClient.DownloadsQueue)[m.selectedRow].Abort()
			return m.Update(UpdateTuiMsg{})
		case key.Matches(msg, constants.Keys.PauseAndPlay):
			(*m.torrentClient.DownloadsQueue)[m.selectedRow].PauseAndPlay()
			return m.Update(msg)
		case key.Matches(msg, constants.Keys.Quit):
			m.torrentClient.Abort()
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Model) View() string {
	return m.table.View() + "\n" + m.help.View()
}

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
		selectedRow:   0,
		torrentClient: torrentClient,
	}
	return m, nil
}
