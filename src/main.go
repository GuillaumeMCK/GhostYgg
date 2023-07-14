package main

import (
	"GhostYgg/src/tui"
	"GhostYgg/src/utils"
	"flag"
	"fmt"
	"github.com/sqweek/dialog"
	"log"
	"os"
	"path/filepath"
)

var (
	torrentFiles   []string
	output         *string
	helpFlag       *bool
	downloadFolder string
)

func init() {
	torrentFiles = []string{}
	// Filter all args that are not torrent files
	for _, arg := range os.Args[1:] {
		if filepath.Ext(arg) == ".torrent" {
			torrentFiles = append(torrentFiles, arg)
		}
	}
	output = flag.String("output", "", "Download directory")
	helpFlag = flag.Bool("help", false, "Show this message")

	// Define how to use the program
	flag.Usage = func() {
		fmt.Printf("GhostYgg - Download torrents\n\n")
		fmt.Printf("Usage: %s file1.client file2.torrent ... [options]\n\n", os.Args[0])
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

	if len(torrentFiles) == 0 {
		// If no file is specified, show the file picker dialog
		if dialog.Message("%s", "No client file specified. Do you want to choose a file?").YesNo() {
			if len(torrentFiles) == 0 {
				// Open file explorer to choose a .client file
				filePath, err := dialog.File().Filter("Torrent files", "torrent").Load()
				if err != nil {
					log.Fatal(err)
				}
				torrentFiles = []string{filePath}
			}
		} else {
			torrentFiles = []string{}
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
	err := tui.StartTea(torrentFiles, downloadFolder)
	if err != nil {
		log.Fatal(err)
	}
}
