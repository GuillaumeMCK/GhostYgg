package tui

import (
	"GhostYgg/src/tui/constants"
	"GhostYgg/src/utils"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

type FilePicker struct {
	input textinput.Model
	size  *utils.Size
	err   string
}

func (m *FilePicker) Init() tea.Cmd {
	m.input.Width = m.size.Width
	return nil
}

func (m *FilePicker) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.(type) {
	case UpdateContainerMsg:
		m.updateWidth()
		return m, nil
	case ClearErrorMsg:
		m.Clear()
		return m, nil
	}
	m.input, cmd = m.input.Update(msg)

	if m.input.Focused() {
		return m, cmd
	}

	return m, cmd
}

func (m *FilePicker) View() string {
	m.updateWidth()
	var s strings.Builder
	s.WriteString(constants.FirstCharFilePicker)
	if !m.input.Focused() {
		return ""
	} else if m.err != "" {
		s.WriteString(constants.ErrorStyle.Render(m.err))
	} else if m.input.Value() == "" {
		s.WriteString("Drag and drop a torrent file here or type the path to a torrent file and press Enter")
	} else {
		s.WriteString("Path: ")
		s.WriteString(m.input.View())
	}
	s.WriteString("\n")
	return s.String()
}

func (m *FilePicker) Focus() tea.Cmd {
	return m.input.Focus()
}

func (m *FilePicker) SetError(err string) {
	m.err = err
	clearErrorAfter(2)
}

func (m *FilePicker) GetValue() string {
	return m.input.Value()
}

func (m *FilePicker) SetValue(value string) {
	m.input.SetValue(value)
}

func (m *FilePicker) Clear() {
	m.input.SetValue("")
	m.err = ""
}

func (m *FilePicker) updateWidth() {
	m.input.Width = m.size.Width
}

func (m *FilePicker) getHeight() int {
	if !m.input.Focused() {
		return 1
	}
	return 2
}

func NewFilePicker(size *utils.Size) *FilePicker {
	input := textinput.New()
	input.PromptStyle = constants.PromptStyle
	input.TextStyle = constants.TextStyle
	input.PlaceholderStyle = constants.PlaceholderStyle
	input.CharLimit = 4096
	input.Width = size.Width

	return &FilePicker{
		input: input,
		size:  size,
	}
}
