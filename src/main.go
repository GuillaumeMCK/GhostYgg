package main

import (
	"GhostYgg/src/utils"
	"flag"
	"fmt"
	"github.com/anacrolix/torrent"
	"github.com/fatih/color"
	"github.com/sqweek/dialog"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var (
	filesFlag []string
	output    *string
	helpFlag  *bool
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
		fmt.Printf("Usage: %s file1.torrent file2.torrent ... \n\n", filepath.Base(os.Args[0]))
		fmt.Printf("Download torrents.\n\n")
		fmt.Printf("Options:\n")
		flag.PrintDefaults()
	}
}

func main() {
	// Parse command-line flags
	flag.Parse()

	// If -help flag is set, show the usage and exit
	if *helpFlag {
		flag.Usage()
		return
	}

	if len(filesFlag) == 0 {
		// If no file is specified, show the file picker dialog
		if dialog.Message("%s", "No torrent file specified. Do you want to choose a file?").YesNo() {
			if len(filesFlag) == 0 {
				// Open file explorer to choose a .torrent file
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
	downloadFolder := *output
	if downloadFolder == "" {
		downloadFolder = defaultDownloadFolder
	}

	// Make sure download folder exists
	err = os.MkdirAll(downloadFolder, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new configuration for the torrent client
	clientConfig := createClientConfig(downloadFolder)

	// Create a new torrent client
	client, err := torrent.NewClient(clientConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Make sure to close the client after use
	defer client.Close()

	// Add the torrent files to the client
	for i, file := range filesFlag {

		// Check if the file already exists in the download folder
		err = utils.Exist(downloadFolder, file)
		if err != nil {
			log.Fatal(err)
		}

		// Add the torrent file to the client
		t, err := client.AddTorrentFromFile(file)
		if err != nil {
			log.Fatal(err)
		}

		// Start downloading the torrent
		t.DownloadAll()

		// Track download progress
		go trackDownloadProgress(t, i)
	}

	// Handle system signals
	handleInterruptSignal()
	handleTerminalResize()

	// Wait for all torrents to finish downloading
	client.WaitAll()
	time.Sleep(1 * time.Second)
	fmt.Printf(color.GreenString("\n\nAll downloads completed. Files saved in %s\n", downloadFolder))

	// Exit the program
	os.Exit(0)
}

func trackDownloadProgress(t *torrent.Torrent, i int) {
	// Wait for the torrent to get info
	<-t.GotInfo()

	down := strings.Repeat(utils.DOWN, i)
	up := strings.Repeat(utils.UP, i)
	percent := 0
	name := t.Info().Name

	// Track download progress
	startTime := time.Now()
	for {
		// Get the percentage of the torrent that is downloaded
		percent = int(t.BytesCompleted() * 100 / t.Info().TotalLength())

		// Calculate download rate in MB/s
		elapsedTime := time.Since(startTime)
		downloadRate := float64(t.BytesCompleted()) / elapsedTime.Seconds() / 1024 / 1024

		// Adjust the name length dynamically based on the console width
		consoleWidth, _, err := terminal.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			log.Fatal(err)
		}
		maxNameLength := consoleWidth - 10
		if len(name) > maxNameLength {
			name = name[:maxNameLength] + "..."
		}
		fmt.Printf("%s\r[%s] %s %s seed:%s leech:%s Rate: %s %s%s",
			down,
			utils.GetDateTime(),
			color.YellowString(utils.FormatBytesProgress(t.BytesCompleted(), t.Info().TotalLength())),
			color.MagentaString(strconv.Itoa(percent)+"%"),
			color.GreenString(strconv.Itoa(t.Stats().ConnectedSeeders)),
			color.RedString(strconv.Itoa(t.Stats().ActivePeers-t.Stats().ConnectedSeeders)),
			color.CyanString("%.2fMB/s", downloadRate),
			name,
			up,
		)

		// If the torrent is fully downloaded, stop tracking progress
		if t.BytesCompleted() == t.Info().TotalLength() {
			break
		}

		time.Sleep(500 * time.Millisecond)
	}

	fmt.Printf("%s\rDownload completed: %s%s\n", down, color.GreenString(name), up)
}

func createClientConfig(downloadFolder string) *torrent.ClientConfig {
	// Create a new configuration for the torrent client and disable http trackers
	clientConfig := torrent.NewDefaultClientConfig()
	clientConfig.DataDir = downloadFolder
	clientConfig.DisableTrackers = false
	clientConfig.Seed = false
	clientConfig.NoUpload = true
	clientConfig.DisableIPv6 = false
	clientConfig.Debug = false
	clientConfig.DisableWebtorrent = true
	return clientConfig
}

func handleInterruptSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			fmt.Println(color.RedString("\nDownload interrupted..."))
			os.Exit(0)
		}
	}()
}

func handleTerminalResize() {
	resize := make(chan os.Signal, 1)
	signal.Notify(resize, syscall.SIGWINCH)
	go func() {
		for range resize {
			<-resize
			utils.ClearScreen()
		}
	}()
}
