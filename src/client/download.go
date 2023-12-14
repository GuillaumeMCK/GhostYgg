package client

import (
	"github.com/GuillaumeMCK/GhostYgg/src/tui/constants"
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
	if d.paused && !d.dropped {
		d.Infos.reset()
		d.SetETA(constants.Paused)
	} else if !d.dropped {
		d.SetETA(constants.Play)
	}
}

func (d *TorrentInfos) Finished() {
	d.finished = true
	d.SetETA(constants.Validated)
}

func (d *TorrentInfos) Abort() {
	d.aborted = true
	if !d.dropped {
		d.SetETA(constants.Cross)
		d.Dropped()
	}
}

func (d *TorrentInfos) Dropped() {
	if !d.dropped {
		d.dropped = true
		d.Infos.reset()
	}
}

func (d *TorrentInfos) Idx() int {
	return d.index
}

func (d *TorrentInfos) IsRunning() bool {
	return !d.aborted && !d.dropped && !d.finished
}
