package tui

import (
	"GhostYgg/src/client"
	"GhostYgg/src/tui/constants"
	"GhostYgg/src/utils"
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

// TUI represents the TUI model
type TUI struct {
	table         *Table
	help          *Help
	header        *Header
	torrentClient *client.Model
	filepicker    *FilePicker
	container     *utils.Size
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
		return m, tea.Batch(cmd, updateContainer(), updateTuiLoop())
	case UpdateContainerMsg:
		m.container.Resize(
			constants.WindowSize.Width,
			constants.WindowSize.Height-m.help.getHeight()-m.header.getHeight())
		if m.filepicker.shown {
			m.filepicker.Update(msg)
		}
		m.table.refresh(m.torrentClient.Torrents)
		m.table.Update(msg)
	case AddTorrentMsg:
		err := m.torrentClient.AddTorrent(msg.Path)
		if err != nil {
			return nil, nil
		}
		m.table.refresh(m.torrentClient.Torrents)
		m.table.Update(msg)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.Keys.Help):
			m.help.Update(msg)
			return m, updateContainer()
		case key.Matches(msg, constants.Keys.Add):
			m.filepicker.Update(msg)
			return m, updateContainer()
		case key.Matches(msg, constants.Keys.Open):
			utils.OpenDirectory(constants.DownloadFolder)
			return m, nil
		case key.Matches(msg, constants.Keys.Delete):
			if len(m.torrentClient.Torrents) == 0 {
				return m, nil
			}
			m.torrentClient.Torrents[m.table.selectedRow()].Abort()
			return m, updateContainer()
		case key.Matches(msg, constants.Keys.PauseAndPlay):
			if len(m.torrentClient.Torrents) == 0 {
				return m, nil
			}
			m.torrentClient.Torrents[m.table.selectedRow()].PauseAndPlay()
			return m, updateContainer()
		case key.Matches(msg, constants.Keys.Quit):
			if m.filepicker.shown {
				m.filepicker.Update(msg)
				return m, nil
			}
			m.torrentClient.Abort()
			return m, tea.Quit
		}
	}
	m.table.Update(msg)
	m.filepicker.Update(msg)
	return m, cmd
}

// View renders the TUI view as a string.
func (m TUI) View() string {
	var s strings.Builder
	s.WriteString(m.header.View())
	if m.filepicker.shown {
		s.WriteString(m.filepicker.View())
	} else {
		s.WriteString(m.table.View())
	}
	s.WriteString("\n" + m.help.View())
	return constants.BaseContainerStyle.Render(s.String())
}

// NewTUI creates a new TUI model.
func NewTUI(torrentFiles []string) (tea.Model, tea.Cmd) {
	torrentClient, _ := client.New(constants.DownloadFolder, torrentFiles)

	err := torrentClient.Start()
	if err != nil {
		panic("error starting torrent client")
	}

	var containerSize = utils.NewSize(80, 24)

	m := TUI{
		table: NewTable(&TableCtx{
			Rows:    make([]client.TorrentInfos, 0),
			Columns: constants.TableColumns,
			Widths:  constants.TableWidths,
			Size:    &containerSize,
		}),
		help:          NewHelp(&containerSize),
		torrentClient: torrentClient,
		filepicker:    NewFilePicker(&containerSize),
		container:     &containerSize,
		header:        NewHeader(),
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
