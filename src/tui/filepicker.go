package tui

import (
	"errors"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
	"time"

	"GhostYgg/src/tui/constants"
	"GhostYgg/src/utils"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
)

// FilePicker represents the TUI file picker model.
type FilePicker struct {
	filepicker   filepicker.Model
	selectedFile string
	shown        bool
	size         *utils.Size
	err          error
}

func (m FilePicker) Init() tea.Cmd {
	return m.filepicker.Init()
}

func (m *FilePicker) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.Keys.Add):
			if m.selectedFile != "" {
				m.shown = false
				return m, tea.Batch(cmd, addTorrent(m.selectedFile))
			}
			m.shown = true
			return m, nil
		case key.Matches(msg, constants.Keys.Quit):
			m.shown = false
			return m, nil
		}
	case clearErrorMsg:
		m.err = nil
	}

	// Did the user select a file?
	if didSelect, path := m.filepicker.DidSelectFile(msg); didSelect {
		m.selectedFile = path
	}

	// Did the user select a disabled file?
	// This is only necessary to display an error to the user.
	if didSelect, path := m.filepicker.DidSelectDisabledFile(msg); didSelect {
		// Let's clear the selectedFile and display an error.
		m.err = errors.New(path + " is not valid.")
		m.selectedFile = ""
		return m, tea.Batch(cmd, clearErrorAfter(2*time.Second))
	}

	m.filepicker, cmd = m.filepicker.Update(msg)
	return m, cmd
}

func (m FilePicker) View() string {
	if !m.shown {
		return ""
	}
	m.updateHeight()
	// print with padding top to 1
	return lipgloss.NewStyle().
		PaddingTop(1).
		Render(m.filepicker.View())
}

func (m FilePicker) updateHeight() {
	m.filepicker.Height = m.size.Height
}

func NewFilePicker(size *utils.Size) *FilePicker {
	fp := filepicker.New()
	fp.Styles = constants.FilePickerStyle
	fp.AllowedTypes = []string{".torrent"}
	fp.CurrentDirectory, _ = utils.GetHomeDir()
	fp.Height = size.Height - 2
	fp.ShowHidden = false
	fp.AutoHeight = false
	fp.Cursor = " "

	return &FilePicker{
		filepicker: fp,
		shown:      false,
		size:       size,
	}
}
