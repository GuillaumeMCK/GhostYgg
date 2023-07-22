package client

import (
	"GhostYgg/src/tui/constants"
	"GhostYgg/src/utils"
	"fmt"
	"github.com/anacrolix/torrent"
	"strconv"
	"time"
)

// Constants
const (
	megabyte           = 1024 * 1024
	torrentSpeedFormat = "%.2fMB/s"
)

// Model represents the torrent client model.
type Model struct {
	Torrents *[]TorrentInfos
	client   *torrent.Client
	files    []string
}

// New creates a new torrent client model.
func New(downloadFolder string, files []string) (*Model, error) {
	// Create a new configuration for the torrent client
	clientConfig := createClientConfig(downloadFolder)
	// Create a new torrent client
	client, err := torrent.NewClient(clientConfig)
	if err != nil {
		return nil, err
	}
	TorrentsInfos := make([]TorrentInfos, 0)
	// Create a new model
	return &Model{Torrents: &TorrentsInfos, client: client, files: files}, nil
}

// Start starts the download process for the client.
func (m *Model) Start() error {
	for _, file := range m.files {
		m.AddTorrent(file)
	}
	return nil
}

// AddTorrent adds a new torrent infos to the model.
func (m *Model) AddTorrent(path string) error {
	var torrentInfos TorrentInfos
	var lenght int = len(*m.Torrents)

	t, err := m.client.AddTorrentFromFile(path)
	if err != nil {
		torrentInfos = defaultTorrentInfos(err.Error(), lenght, path)
		torrentInfos.Abort()
	} else {
		torrentInfos = defaultTorrentInfos(t.Info().Name, lenght, path)
	}
	// Add the torrent infos to the model
	*m.Torrents = append(*m.Torrents, torrentInfos)

	// Start downloading the client
	t.DownloadAll()

	// Start tracking the torrent
	go m.trackTorrent(t, lenght)

	return nil
}

// trackTorrent tracks the download progress of a client.
func (m *Model) trackTorrent(t *torrent.Torrent, index int) {
	<-t.GotInfo()

	name := t.Info().Name
	startTime := time.Now()
	startSize := t.BytesCompleted()

	for {
		torrentInfos := (*m.Torrents)[index]
		bytesCompleted := t.BytesCompleted()
		totalLength := t.Info().TotalLength()
		elapsedTime := time.Since(startTime)

		if torrentInfos.aborted || torrentInfos.paused {
			if !torrentInfos.dropped {
				t.Drop()
				torrentInfos.dropped = true
			}
		} else if bytesCompleted >= totalLength {
			torrentInfos.finished = true
			torrentInfos.SetETA(constants.Validated)
		} else if torrentInfos.dropped && !torrentInfos.paused {
			t, _ = m.client.AddTorrentFromFile(torrentInfos.path)
			t.DownloadAll()
			startSize = t.BytesCompleted()
			startTime = time.Now()
			torrentInfos.dropped = false
		}

		if bytesCompleted < totalLength &&
			!torrentInfos.finished &&
			!torrentInfos.aborted &&
			!torrentInfos.paused {
			remainingBytes := totalLength - bytesCompleted
			downloadRate := calculateDownloadRate(bytesCompleted, startSize, elapsedTime)
			torrentInfos.Infos = Infos{
				Name:         name,
				Progress:     utils.FormatBytesProgress(bytesCompleted, totalLength),
				Seeders:      strconv.Itoa(t.Stats().ConnectedSeeders),
				Leechers:     strconv.Itoa(t.Stats().ActivePeers - t.Stats().ConnectedSeeders),
				Torrentspeed: fmt.Sprintf(torrentSpeedFormat, downloadRate),
				ETA:          utils.FormatDuration(calculateETA(remainingBytes, downloadRate)),
			}
		}

		// Write the torrent infos to the model
		(*m.Torrents)[index] = torrentInfos

		if torrentInfos.finished || torrentInfos.aborted {
			break
		}
		time.Sleep(150 * time.Millisecond)
	}
}

// calculateDownloadRate calculates the download rate in MB/s.
func calculateDownloadRate(bytesCompleted, startSize int64, elapsedTime time.Duration) float64 {
	return float64(bytesCompleted-startSize) / elapsedTime.Seconds() / megabyte
}

// calculateETA calculates the estimated time of arrival for the download completion.
func calculateETA(remainingBytes int64, downloadRate float64) time.Duration {
	if downloadRate > 0 {
		return time.Duration(float64(remainingBytes)/downloadRate) / megabyte
	}
	return time.Duration(0)
}

// Abort aborts all Torrents in the client model.
func (m *Model) Abort() {
	for _, torrentInfos := range *m.Torrents {
		torrentInfos.Abort()
	}
}

// createClientConfig creates a new configuration for the torrent client.
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
