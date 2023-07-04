package main

import (
	"GhostYgg/src/utils"
	"flag"
	"fmt"
	"github.com/anacrolix/torrent"
	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
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

	// If no file is specified, show the usage and exit
	if len(filesFlag) == 0 {
		flag.Usage()
		return
	}

	// Determine default download folder
	defaultDownloadFolder, err := getDefaultDownloadFolder()
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

	// Handle system interrupt signal (ctrl+c)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			fmt.Println("\nDownload interrupted...")
			os.Exit(1)
		}
	}()

	// Wait for all torrents to finish downloading
	client.WaitAll()
}

func trackDownloadProgress(t *torrent.Torrent, i int) {
	// Wait for the torrent to get info
	<-t.GotInfo()

	down := strings.Repeat(utils.DOWN, i)
	up := strings.Repeat(utils.UP, i)
	percent := 0
	// If the name is too long, cap it
	name := t.Info().Name
	if len(name) > 50 {
		name = name[:50] + "..."
	}

	// Track download progress
	for {
		// Get the percentage of the torrent that is downloaded
		percent = int(t.BytesCompleted() * 100 / t.Info().TotalLength())

		fmt.Printf("%s\r[%s] status: %s/%s %s seeders:%s name:%s%s",
			down,
			utils.GetDateTime(),
			color.CyanString(utils.ByteSuffixes(t.BytesCompleted())),
			utils.ByteSuffixes(t.Info().TotalLength()),
			color.MagentaString(strconv.Itoa(percent)+"%"),
			color.GreenString(strconv.Itoa(t.Stats().ConnectedSeeders)),
			name,
			up,
		)

		// If the torrent is fully downloaded, stop tracking progress
		if t.BytesCompleted() == t.Info().TotalLength() {
			break
		}

		time.Sleep(250 * time.Millisecond)
	}

	fmt.Println(color.GreenString("\n\nDownload completed: %s", t.Info().Name))
}

func getDefaultDownloadFolder() (string, error) {
	homeDir, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	downloadFolder := ""
	switch {
	case os.Getenv("HOME") != "":
		downloadFolder = os.Getenv("HOME") + "/Downloads"
	case os.Getenv("USERPROFILE") != "":
		downloadFolder = os.Getenv("USERPROFILE") + "\\Downloads"
	default:
		downloadFolder = homeDir
	}

	return downloadFolder, nil
}

func createClientConfig(downloadFolder string) *torrent.ClientConfig {
	// Create a new configuration for the torrent client
	clientConfig := torrent.NewDefaultClientConfig()
	clientConfig.DataDir = downloadFolder
	clientConfig.DisableTrackers = false
	clientConfig.Seed = false
	clientConfig.NoUpload = true
	clientConfig.DisableIPv6 = false
	clientConfig.Debug = false

	return clientConfig
}
