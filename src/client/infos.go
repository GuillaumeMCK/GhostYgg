package client

type Infos struct {
	Name          string
	Progress      string
	Seeders       string
	Leeches       string
	DownloadSpeed string
	ETA           string
}

func (i Infos) ToStringArray() [6]string {
	return [6]string{i.Name, i.Progress, i.Seeders, i.Leeches, i.DownloadSpeed, i.ETA}
}

func (i *Infos) reset() {
	i.Seeders = "0"
	i.Leeches = "0"
	i.DownloadSpeed = "0.00MB/s"
}

func defaultInfos(name string) Infos {
	return Infos{
		Name:          name,
		Progress:      "",
		Seeders:       "0",
		Leeches:       "0",
		DownloadSpeed: "0.00MB/s",
		ETA:           "00:00:00",
	}
}
