package constants

import (
	tea "github.com/charmbracelet/bubbletea"
)

var (
	// P the current tea program
	P *tea.Program
	// WindowSize store the size of the terminal window
	WindowSize tea.WindowSizeMsg = tea.WindowSizeMsg{Width: 80, Height: 24}
	// DownloadFolder is the folder where the torrents will be downloaded
	DownloadFolder string
	// TableColumns are the columns of the table
	TableColumns = [6]string{"Name", "Progress", "Seeders " + UpArrow, "Leeches " + DownArrow, "Download Speed", "ETA"}
	// TableWidths are the widths of the table columns
	TableWidths = [6]float32{0.4, 0.15, 0.1, 0.1, 0.15, 0.1}
	// Exceptions
	ErrNotTorrentFile = "Invalid file type. Please select a .torrent file."
	ErrFileNotFound   = "File not found."
)
