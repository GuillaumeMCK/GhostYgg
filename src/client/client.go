package client

import (
	"GhostYgg/src/utils"
	"fmt"
	"github.com/anacrolix/log"
	"github.com/anacrolix/torrent"
	"strconv"
	"time"
)

type Infos struct {
	Name          string
	Progress      string
	Seeders       string
	Leeches       string
	DownloadSpeed string
	ETA           string
}

func (i Infos) ToStringArray() []string {
	return []string{i.Name, i.Progress, i.Seeders, i.Leeches, i.DownloadSpeed, i.ETA}
}

type DownloadInfos struct {
	Infos    Infos
	Index    int
	Finished bool
}

type Commands struct {
	UpdateRow func([]string, ...int)
	Exit      func(int)
}

func Start(downloadFolder string, files []string, commands Commands) {
	// Create a new configuration for the torrent client
	clientConfig := createClientConfig(downloadFolder)

	// Create a new torrent client
	client, err := torrent.NewClient(clientConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Make sure to close the client after use
	defer client.Close()

	// Add the client files to the client
	for i, file := range files {
		// Add the client file to the client
		t, err := client.AddTorrentFromFile(file)
		if err != nil {
			log.Fatal(err)
		}

		// Create DownloadInfos object
		downloadInfo := DownloadInfos{
			Infos: Infos{
				Name:          t.Info().Name,
				Progress:      "0",
				Seeders:       "0",
				Leeches:       "0",
				DownloadSpeed: "0",
				ETA:           "00:00:00",
			},
			Index:    i,
			Finished: false,
		}

		// Start downloading the client
		t.DownloadAll()
		// Add the client to the table
		commands.UpdateRow(downloadInfo.Infos.ToStringArray())
		// Track download progress
		go trackDownloadProgress(t, downloadInfo, commands)
	}

	// Wait for all torrents to finish downloading
	client.WaitAll()

	// Exit the program
	commands.Exit(0)
}

// trackDownloadProgress tracks the download progress of a client
func trackDownloadProgress(t *torrent.Torrent, downloadInfo DownloadInfos, commands Commands) {
	// Wait for the client to get info
	<-t.GotInfo()

	// Define variables
	name := t.Info().Name
	startTime := time.Now()
	eta := int64(0)
	startSize := t.BytesCompleted()

	for {
		// If the client is still downloading
		if t.BytesCompleted() < t.Info().TotalLength() {
			elapsedTime := time.Since(startTime)
			downloadRate := float64(t.BytesCompleted()-startSize) / elapsedTime.Seconds()
			remainingBytes := t.Info().TotalLength() - t.BytesCompleted()
			if downloadRate > 0 {
				eta = remainingBytes / int64(downloadRate*1024*1024) // Remaining bytes / download rate in MB/s
			}
			downloadInfo.Infos = Infos{
				Name:          name,
				Progress:      utils.FormatBytesProgress(t.BytesCompleted(), t.Info().TotalLength()),
				Seeders:       strconv.Itoa(t.Stats().ConnectedSeeders),
				Leeches:       strconv.Itoa(t.Stats().ActivePeers - t.Stats().ConnectedSeeders),
				DownloadSpeed: fmt.Sprintf("%.2fMB/s", downloadRate),
				ETA:           utils.FormatDuration(time.Duration(eta) * time.Second),
			}

		} else {
			// Update Infos attribute on download completion
			downloadInfo.Infos = Infos{
				Name:          name,
				Progress:      utils.FormatBytesProgress(t.BytesCompleted(), t.Info().TotalLength()),
				Seeders:       "0",
				Leeches:       "0",
				DownloadSpeed: "0",
				ETA:           "00:00:00",
			}
			// Mark download as finished
			downloadInfo.Finished = true
		}
		// Update the client
		commands.UpdateRow(downloadInfo.Infos.ToStringArray(), downloadInfo.Index)

		time.Sleep(250 * time.Millisecond)
	}
}

// createClientConfig creates a new configuration for the torrent client
func createClientConfig(downloadFolder string) *torrent.ClientConfig {
	clientConfig := torrent.NewDefaultClientConfig()
	clientConfig.DataDir = downloadFolder
	clientConfig.DisableTrackers = false
	clientConfig.Seed = false
	clientConfig.NoUpload = true
	clientConfig.DisableIPv6 = false
	clientConfig.Debug = false
	clientConfig.DisableWebtorrent = true
	clientConfig.DisableWebseeds = true
	return clientConfig
}
