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
		key.WithKeys("o", "enter"),
		key.WithHelp("o/enter", "open download folder"),
	),
	Add: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "add torrent to client"),
	),
	Delete: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete torrent from client"),
	),
	PauseAndPlay: key.NewBinding(
		key.WithKeys("space"),
		key.WithHelp("space", "pause/play a torrent"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q", "esc"),
		key.WithHelp("ctrl+c/q/esc", "quit"),
	),
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "move the cursor up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "move the cursor down"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Open, k.Add, k.Delete, k.PauseAndPlay, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Open, k.Add, k.Delete, k.PauseAndPlay}, // first column
		{k.Up, k.Down, k.Help, k.Quit},            // second column
	}
}
