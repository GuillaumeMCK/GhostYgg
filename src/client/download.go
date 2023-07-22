package client

import (
	"GhostYgg/src/tui/constants"
)

type TorrentInfos struct {
	Infos    Infos
	index    int
	finished bool
	paused   bool
	aborted  bool
	dropped  bool
	path     string
}

func defaultTorrentInfos(name string, index int, path string) TorrentInfos {
	return TorrentInfos{
		Infos:    defaultInfos(name),
		index:    index,
		finished: false,
		paused:   false,
		aborted:  false,
		dropped:  false,
		path:     path,
	}
}
func (d *TorrentInfos) SetETA(eta string) {
	d.Infos.ETA = eta
}

func (d *TorrentInfos) PauseAndPlay() {
	d.paused = !d.paused
	if d.paused && !d.finished && !d.aborted {
		d.SetETA(constants.Paused)
	}
}

func (d *TorrentInfos) Abort() {
	d.aborted = true
	if !d.finished {
		d.SetETA(constants.Cross)
	}
}

func (d *TorrentInfos) Idx() int {
	return d.index
}
