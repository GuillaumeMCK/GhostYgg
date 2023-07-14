package main

import (
	"GhostYgg/src/client"
	"GhostYgg/src/tui"
	"GhostYgg/src/utils"
	"flag"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/sqweek/dialog"
	"log"
	"os"
	"path/filepath"
)

// tui args
var (
	filesFlag []string
	output    *string
	helpFlag  *bool
)

// table
var (
	tableCtx = &tui.TableCtx{
		Columns: []string{"Name", "Progress", "Seeders", "Leeches", "Download", "ETA"}, // up arrow special char : ↑ | down arrow special char : ↓ | up and down arrow like download special char : ↕
		Rows:    [][]string{},
		Widths:  []float32{0.4, 0.1, 0.1, 0.1, 0.1, 0.1},
	}
	downloadFolder string
)

func init() {
	filesFlag = []string{}
	// Filter all args that are not torrent files
	for _, arg := range os.Args[1:] {
		if filepath.Ext(arg) == ".torrent" {
			filesFlag = append(filesFlag, arg)
		}
	}
	output = flag.String("output", "", "Download directory")
	helpFlag = flag.Bool("help", false, "Show this message")

	// Define how to use the program
	flag.Usage = func() {
		fmt.Printf("GhostYgg - Download torrents\n\n")
		fmt.Printf("Usage: %s file1.client file2.torrent ... \n\n", filepath.Base(os.Args[0]))
		fmt.Printf("Download torrents.\n\n")
		fmt.Printf("Options:\n")
		flag.PrintDefaults()
	}

	// Parse command-line flags
	flag.Parse()

	// If -help flag is set, show the usage and exit
	if *helpFlag {
		flag.Usage()
		return
	}

	if len(filesFlag) == 0 {
		// If no file is specified, show the file picker dialog
		if dialog.Message("%s", "No client file specified. Do you want to choose a file?").YesNo() {
			if len(filesFlag) == 0 {
				// Open file explorer to choose a .client file
				filePath, err := dialog.File().Filter("Torrent files", "torrent").Load()
				if err != nil {
					log.Fatal(err)
				}
				filesFlag = []string{filePath}
			}
		} else {
			os.Exit(0)
		}
	}

	// Determine default download folder
	defaultDownloadFolder, err := utils.GetDefaultDownloadFolder()
	if err != nil {
		log.Fatal(err)
	}

	// Use the value of -download-folder flag, or default download folder
	downloadFolder = *output
	if downloadFolder == "" {
		downloadFolder = defaultDownloadFolder
	}

	// Make sure download folder exists
	err = os.MkdirAll(downloadFolder, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// define rows in tableCtx
	tableCtx.Rows = make([][]string, 0)
	// Create a new table with the default configuration
	t := tui.Table(tableCtx)
	// Run the client
	go client.Start(downloadFolder, filesFlag, client.Commands{
		UpdateRow: func(row []string, index ...int) {
			if len(index) == 0 {
				tableCtx.Rows = append(tableCtx.Rows, row)
			} else {
				tableCtx.Rows[index[0]] = row
			}
			t.Update(tableCtx)
		},
		Exit: func(code int) {
			t.Exit()
		},
	})

	// Start the Bubble Tea program
	_, err := tea.NewProgram(t,
		tea.WithAltScreen(),
		tea.WithMouseAllMotion(),
		tea.WithMouseCellMotion(),
	).Run()

	if err != nil {
		fmt.Println("Oh no!", err)
	}

	_, _ = fmt.Scanln() // Ignore the input
}

// [2023-07-14 22:44:38 +0200  NIL] requester 8: error doing webseed request {103 {0 16384}}: reading "connect.joinpeers.org" at "bytes=54001664-54018047": Get "connect.joinpeers.org": unsupported protocol scheme "" [github.com/anacrolix/torrent webseed-peer.go:106]
//[2023-07-14 22:44:38 +0200  NIL] requester 9: error doing webseed request {103 {0 16384}}: reading "connect.joinpeers.org" at "bytes=54001664-54018047": Get "connect.joinpeers.org": unsupported protocol scheme "" [github.com/anacrolix/torrent webseed-peer.go:106]
//[2023-07-14 22:44:38 +0200  NIL] requester 13: error doing webseed request {103 {0 16384}}: reading "connect.joinpeers.org" at "bytes=54001664-54018047": Get "connect.joinpeers.org": unsupported protocol scheme "" [github.com/anacrolix/torrent webseed-peer.go:106]
//[2023-07-14 22:44:38 +0200  NIL] requester 12: error doing webseed request {103 {0 16384}}: reading "connect.joinpeers.org" at "bytes=54001664-54018047": Get "connect.joinpeers.org": unsupported protocol scheme "" [github.com/anacrolix/torrent webseed-peer.go:106]
//[2023-07-14 22:44:38 +0200  NIL] requester 14: error doing webseed request {103 {0 16384}}: reading "connect.joinpeers.org" at "bytes=54001664-54018047": Get "connect.joinpeers.org": unsupported protocol scheme "" [github.com/anacrolix/torrent webseed-peer.go:106]
//[2023-07-14 22:44:38 +0200  NIL] requester 10: error doing webseed request {103 {0 16384}}: reading "connect.joinpeers.org" at "bytes=54001664-54018047": Get "connect.joinpeers.org": unsupported protocol scheme "" [github.com/anacrolix/torrent webseed-peer.go:106]
//[2023-07-14 22:44:38 +0200  NIL] requester 15: error doing webseed request {103 {0 16384}}: reading "connect.joinpeers.org" at "bytes=54001664-54018047": Get "connect.joinpeers.org": unsupported protocol scheme "" [github.com/anacrolix/torrent webseed-peer.go:106]
//[2023-07-14 22:44:40 +0200  WRN] UPnP device at 192.168.1.73: mapping internal UDP port 42069: error: AddPortMapping: 500 Internal Server Error [github.com/anacrolix/torrent portfwd.go:17]
//[2023-07-14 22:44:40 +0200  WRN] UPnP device at 192.168.1.73: mapping internal TCP port 42069: error: AddPortMapping: 500 Internal Server Error [github.com/anacrolix/torrent portfwd.go:17]
