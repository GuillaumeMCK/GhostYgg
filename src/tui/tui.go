package tui

import (
	"fmt"
	"github.com/GuillaumeMCK/GhostYgg/src/client"
	"github.com/GuillaumeMCK/GhostYgg/src/tui/constants"
	"github.com/GuillaumeMCK/GhostYgg/src/utils"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// TUI represents the TUI model
type TUI struct {
	table         *Table
	help          *Help
	header        *Header
	torrentClient *client.Model
	filePicker    *FilePicker
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
			constants.WindowSize.Height-m.header.getHeight()-m.help.getHeight()-m.filePicker.getHeight())
		m.table.refresh(m.torrentClient.Torrents)
		m.table.Update(msg)
		return m, nil
	case AddTorrentMsg:
		err := m.torrentClient.AddTorrent(msg.Path)
		if err != nil {
			m.filePicker.SetError(err.Error())
			return m, tea.Batch(updateContainer())
		}
		m.table.refresh(m.torrentClient.Torrents)
		m.filePicker.Clear()
		m.filePicker.input.Blur()
		return m, tea.Batch(updateContainer())
	}
	if m.filePicker.input.Focused() {
		return m.handleFilePickerInput(msg)
	} else {
		return m.handleKeyInput(msg)
	}
}

// handleFilePickerInput handles the messages when the file picker is focused.
func (m *TUI) handleFilePickerInput(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.Keys.Enter):
			value := m.filePicker.input.Value()
			if value != "" {
				value = strings.ReplaceAll(value, "\\", "")
				filePath := filepath.Clean(strings.TrimRight(value, " "))
				if filepath.Ext(filePath) != ".torrent" {
					m.filePicker.SetError(constants.ErrNotTorrentFile)
					return m, nil
				} else if !utils.Exist(filePath) {
					m.filePicker.SetError(constants.ErrFileNotFound)
					return m, nil
				}
				return m, tea.Batch(addTorrent(filePath), updateContainer())
			}
		case key.Matches(msg, constants.Keys.Exit):
			m.filePicker.input.Blur()
			m.filePicker.Clear()
			return m, updateContainer()
		}
	}
	m.filePicker.Update(msg)
	return m, nil
}

// handleKeyInput handles the messages when the file picker is not focused.
func (m *TUI) handleKeyInput(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.Keys.Add):
			return m, m.filePicker.input.Focus()
		case key.Matches(msg, constants.Keys.Help):
			m.help.Update(msg)
			return m, updateContainer()
		case key.Matches(msg, constants.Keys.Add):
			m.filePicker.Focus()
			return m, updateContainer()
		case key.Matches(msg, constants.Keys.Open):
			if err := utils.OpenDirectory(constants.DownloadFolder); err != nil {
				return m, nil
			}
			return m, nil
		case key.Matches(msg, constants.Keys.Delete):
			if len(m.torrentClient.Torrents) == 0 {
				return m, nil
			}
			m.torrentClient.Torrents[m.table.selectedRow()].Abort()
			m.table.refresh(m.torrentClient.Torrents)
			return m, updateContainer()
		case key.Matches(msg, constants.Keys.PauseAndPlay):
			if len(m.torrentClient.Torrents) == 0 {
				return m, nil
			}
			torrent := m.torrentClient.Torrents[m.table.selectedRow()]
			if torrent.IsRunning() {
				torrent.PauseAndPlay()
			}
			return m, updateContainer()
		case key.Matches(msg, constants.Keys.Exit):
			if m.filePicker.input.Focused() {
				m.filePicker.input.Blur()
				return m, updateContainer()
			}
			m.torrentClient.Abort()
			return m, tea.Quit
		}
	}
	m.table.Update(msg)
	return m, nil
}

// View renders the TUI view as a string.
func (m TUI) View() string {
	var s strings.Builder
	s.WriteString(m.header.View())
	s.WriteString(m.help.View())
	s.WriteString(m.filePicker.View())
	s.WriteString(m.table.View())
	return s.String()
}

// NewTUI creates a new TUI model.
func NewTUI(torrentFiles []string) (tea.Model, tea.Cmd) {
	torrentClient, err := client.New(constants.DownloadFolder, torrentFiles)
	if err != nil {
		panic("Error creating torrent client")
	}
	err = torrentClient.Start()
	if err != nil {
		panic("Error starting torrent client")
	}

	containerSize := utils.NewSize(80, 24)

	m := TUI{
		table:         NewTable(&TableCtx{Rows: make([]client.TorrentInfos, 0), Columns: constants.TableColumns, Widths: constants.TableWidths, Size: &containerSize}),
		help:          NewHelp(&containerSize),
		torrentClient: torrentClient,
		filePicker:    NewFilePicker(&containerSize),
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
		return fmt.Errorf("Error creating TUI: %v", err)
	}

	constants.P = tea.NewProgram(m, tea.WithAltScreen())
	if _, err := constants.P.Run(); err != nil {
		return fmt.Errorf("Error running program: %v", err)
	}

	return nil
}
