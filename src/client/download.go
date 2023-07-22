package client

import "GhostYgg/src/tui/constants"

type DownloadInfos struct {
	Infos    Infos
	index    int
	finished bool
	paused   bool
	aborted  bool
}

func defaultDownloadInfos(name string, index int) DownloadInfos {
	return DownloadInfos{
		Infos:    defaultInfos(name),
		index:    index,
		finished: false,
		paused:   false,
		aborted:  false,
	}
}
func (d *DownloadInfos) SetETA(eta string) {
	d.Infos.ETA = eta
}

func (d *DownloadInfos) PauseAndPlay() {
	d.paused = !d.paused
	if d.paused && !d.finished && !d.aborted {
		d.SetETA(constants.Paused)
	}
}

func (d *DownloadInfos) Abort() {
	d.aborted = true
	if !d.finished {
		d.SetETA(constants.Cross)
	}
}

func (d *DownloadInfos) Index() int {
	return d.index
}
