package constants

import "github.com/charmbracelet/bubbles/key"

// KeyMap is a collection of key bindings
type KeyMap struct {
	Open         key.Binding
	Add          key.Binding
	Delete       key.Binding
	Up           key.Binding
	Down         key.Binding
	PauseAndPlay key.Binding
	Quit         key.Binding
	Help         key.Binding
}

// Keys reusable key mappings shared across models
var Keys = KeyMap{
	Open: key.NewBinding(
		key.WithKeys("o"),
		key.WithHelp("[o]", "open folder"),
	),
	Add: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("[a]", "add torrent"),
	),
	Delete: key.NewBinding(
		key.WithKeys("backspace"),
		key.WithHelp("[backspace]", "delete"),
	),
	PauseAndPlay: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("[ ]", "pause/play"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("[ctrl+c] [q]", "quit"),
	),
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("[↑]", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("[↓]", "move down"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("[?]", "help"),
	),
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit, k.Add, k.Delete, k.PauseAndPlay}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Open, k.Add, k.Delete, k.PauseAndPlay}, // first column
		{k.Up, k.Down, k.Help, k.Quit},            // second column
	}
}
