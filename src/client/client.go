package client

import (
	"GhostYgg/src/tui/constants"
	"GhostYgg/src/utils"
	"fmt"
	"github.com/anacrolix/log"
	"github.com/anacrolix/torrent"
	"net"
	"strconv"
	"sync"
	"time"
)

// Constants
const (
	megabyte           = 1024 * 1024
	torrentSpeedFormat = "%.2fMB/s"
)

// Model represents the torrent client model.
type Model struct {
	Torrents []TorrentInfos  // Slice to store information about each torrent being downloaded.
	client   *torrent.Client // The underlying torrent client used to manage torrents.
	files    []string        // Paths to the torrent files to be downloaded.
	mux      sync.Mutex      // Mutex to protect concurrent access to Torrents slice.
}

// New creates a new torrent client model.
func New(downloadFolder string, files []string) (*Model, error) {
	// Create a new configuration for the torrent client.
	clientConfig := createClientConfig(downloadFolder)
	// Create a new torrent client.
	client, err := torrent.NewClient(clientConfig)
	if err != nil {
		return nil, err
	}
	TorrentsInfos := make([]TorrentInfos, 0)
	// Create a new model.
	return &Model{Torrents: TorrentsInfos, client: client, files: files}, nil
}

// AddTorrent adds a new torrent info to the model.
func (m *Model) AddTorrent(path string) error {
	var torrentInfos TorrentInfos

	m.mux.Lock()         // Lock the mutex before modifying the Torrents slice.
	defer m.mux.Unlock() // Ensure we unlock the mutex even if there's a panic.

	length := len(m.Torrents)

	t, err := m.client.AddTorrentFromFile(path)
	if err != nil {
		torrentInfos = defaultTorrentInfos(err.Error(), length, path)
		torrentInfos.Abort()
	} else {
		torrentInfos = defaultTorrentInfos(t.Info().Name, length, path)
	}
	// Add the torrent info to the model.
	m.Torrents = append(m.Torrents, torrentInfos)

	// Start downloading the torrent.
	t.DownloadAll()

	// Start tracking the torrent in a separate goroutine.
	go m.trackTorrent(t, length)

	return nil
}

// Start starts the download process for the client.
func (m *Model) Start() error {
	for _, path := range m.files {
		err := m.AddTorrent(path)
		if err != nil {
			return err
		}
	}
	return nil
}

// trackTorrent tracks the download progress of a torrent.
func (m *Model) trackTorrent(t *torrent.Torrent, index int) {
	<-t.GotInfo()

	name := t.Info().Name
	startTime := time.Now()
	startSize := t.BytesCompleted()

	for {
		m.mux.Lock() // Lock the mutex before accessing the Torrents slice.
		torrentInfos := &m.Torrents[index]
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
		} else if bytesCompleted < totalLength {
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

		m.mux.Unlock() // Unlock the mutex after accessing the Torrents slice.

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
	for _, torrentInfos := range m.Torrents {
		torrentInfos.Abort()
	}
}

// getAvailablePort returns an available port by listening on a random port and extracting the chosen port.
func getAvailablePort() (int, error) {
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}
	defer listener.Close()

	_, portString, err := net.SplitHostPort(listener.Addr().String())
	if err != nil {
		return 0, err
	}

	port, err := strconv.Atoi(portString)
	if err != nil {
		return 0, err
	}

	return port, nil
}

// createClientConfig creates a new configuration for the torrent client.
func createClientConfig(downloadFolder string) *torrent.ClientConfig {
	port, err := getAvailablePort()
	if err != nil {
		return nil
	}

	clientConfig := torrent.NewDefaultClientConfig()
	clientConfig.ListenPort = port
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
	clientConfig.Logger.SetHandlers(log.DiscardHandler)
	return clientConfig
}
