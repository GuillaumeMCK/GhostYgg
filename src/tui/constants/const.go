package constants

import (
	"GhostYgg/src/tui/table"
	tea "github.com/charmbracelet/bubbletea"
)

var (
	// P the current tea program
	P *tea.Program
	// WindowSize store the size of the terminal window
	WindowSize tea.WindowSizeMsg
	// HelpHeight is the height of the help context menu
	HelpHeight = 2
	// DownloadFolder is the folder where the torrents will be downloaded
	DownloadFolder string
	// TorrentFiles is the list of torrent files to download
	TorrentFiles []string
	// TableCtx is the context for the table
	TableCtx = &table.TableCtx{
		Columns: [6]string{"Name", "Progress", "Seeders " + UpArrow, "Leeches " + DownArrow, "Download Speed", "ETA"},
		Widths:  [6]float32{0.4, 0.15, 0.1, 0.1, 0.15, 0.1},
		Rows:    [][6]string{},
	}
)
