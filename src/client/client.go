package client

import (
	"GhostYgg/src/tui/constants"
	"GhostYgg/src/utils"
	"fmt"
	"github.com/anacrolix/torrent"
	"log"
	"strconv"
	"time"
)

type Model struct {
	DownloadsInfos []*DownloadInfos
	client         *torrent.Client
}

func New(downloadFolder string, files []string) Model {
	// Create a new configuration for the torrent client
	clientConfig := createClientConfig(downloadFolder)
	// Create a new torrent client
	client, err := torrent.NewClient(clientConfig)
	downloadsInfos := make([]*DownloadInfos, 0)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure to close the client after use
	defer client.Close()
	// Add the client files to the client
	for i, file := range files {
		var downloadInfo DownloadInfos
		t, err := client.AddTorrentFromFile(file)
		if err != nil {
			downloadInfo = defaultDownloadInfos(err.Error(), i)
			downloadInfo.SetETA(constants.Cross)
			fmt.Println(err)
		} else {
			downloadInfo = defaultDownloadInfos(t.Info().Name, i)
		}

		// Start downloading the client
		t.DownloadAll()
		// Add the client to the constants
		//commands.UpdateTorrentInfo(&downloadInfo)
		downloadsInfos = append(downloadsInfos, &downloadInfo)
		// If the download failed, skip the rest
		if err != nil {
			continue // Skip the rest
		}
		// Track download progress
		go trackDownload(t, &downloadInfo)
	}
	return Model{DownloadsInfos: downloadsInfos, client: client}
}

// trackDownload tracks the download progress of a client
func trackDownload(t *torrent.Torrent, downloadInfo *DownloadInfos) {
	// Wait for the client to get info
	<-t.GotInfo()
	// Define variables
	name := t.Info().Name
	startTime := time.Now()
	startSize := t.BytesCompleted()
	for {
		bytesCompleted := t.BytesCompleted()
		totalLength := t.Info().TotalLength()
		elapsedTime := time.Since(startTime)

		if bytesCompleted < totalLength {
			remainingBytes := totalLength - bytesCompleted
			downloadRate := calculateDownloadRate(bytesCompleted, startSize, elapsedTime)
			downloadInfo.Infos = Infos{
				Name:          name,
				Progress:      utils.FormatBytesProgress(bytesCompleted, totalLength),
				Seeders:       strconv.Itoa(t.Stats().ConnectedSeeders),
				Leeches:       strconv.Itoa(t.Stats().ActivePeers - t.Stats().ConnectedSeeders),
				DownloadSpeed: fmt.Sprintf("%.2fMB/s", downloadRate),
				ETA:           utils.FormatDuration(calculateETA(remainingBytes, downloadRate)),
			}
		} else {
			downloadInfo.finished = true
			downloadInfo.SetETA(constants.Validated)
		}

		if downloadInfo.aborted {
			downloadInfo.SetETA(constants.Cross)
			t.Drop()
		} else if downloadInfo.paused {
			t.DisallowDataDownload()
			downloadInfo.SetETA(constants.Paused)
		} else {
			t.AllowDataDownload()
		}

		//commands.UpdateTorrentInfo(downloadInfo)

		//println(downloadInfo.Infos.Name + " " + downloadInfo.Infos.Progress + " " + downloadInfo.Infos.Seeders + " " + downloadInfo.Infos.Leeches + " " + downloadInfo.Infos.DownloadSpeed + " " + downloadInfo.Infos.ETA)
		//// status
		//println("Status finished: " + strconv.FormatBool(downloadInfo.finished) + " aborted: " + strconv.FormatBool(downloadInfo.aborted) + " paused: " + strconv.FormatBool(downloadInfo.paused))

		if downloadInfo.finished || downloadInfo.aborted {
			break
		}
		//time.Sleep(150 * time.Millisecond)
	}
}

func calculateDownloadRate(bytesCompleted, startSize int64, elapsedTime time.Duration) float64 {
	return float64(bytesCompleted-startSize) / elapsedTime.Seconds() / 1024 / 1024
}

func calculateETA(remainingBytes int64, downloadRate float64) time.Duration {
	if downloadRate > 0 {
		return time.Duration(int64(float64(remainingBytes) / downloadRate / 1024 / 1024))
	}
	return time.Duration(0)
}

// Abort all downloads
func (m *Model) Abort() {
	for _, downloadInfo := range m.DownloadsInfos {
		downloadInfo.Abort()
	}
}

// createClientConfig creates a new configuration for the torrent client
func createClientConfig(downloadFolder string) *torrent.ClientConfig {
	clientConfig := torrent.NewDefaultClientConfig()
	clientConfig.DataDir = downloadFolder
	clientConfig.DisableTrackers = false
	clientConfig.Seed = false
	clientConfig.NoUpload = true
	clientConfig.DisableIPv6 = true
	clientConfig.Debug = false
	clientConfig.DisableWebtorrent = true
	clientConfig.DisableWebseeds = true
	clientConfig.DisableAcceptRateLimiting = true
	clientConfig.NoDefaultPortForwarding = true
	return clientConfig
}
