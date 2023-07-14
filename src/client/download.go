package client

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
}

func (d *DownloadInfos) Abort() {
	d.aborted = true
}

func (d *DownloadInfos) Index() int {
	return d.index
}
