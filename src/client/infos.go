package client

type Infos struct {
	Name          string
	Progress      string
	Seeders       string
	Leechers      string
	DownloadSpeed string
	ETA           string
}

func (i *Infos) reset() {
	i.Seeders = "0"
	i.Leechers = "0"
	i.DownloadSpeed = "0.00MB/s"
}

func defaultInfos(name string) Infos {
	return Infos{
		Name:          name,
		Progress:      "",
		Seeders:       "0",
		Leechers:      "0",
		DownloadSpeed: "0.00MB/s",
		ETA:           "00:00:00",
	}
}
