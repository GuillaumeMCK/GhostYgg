package client

type Infos struct {
	Name         string
	Progress     string
	Seeders      string
	Leechers     string
	Torrentspeed string
	ETA          string
}

func (i *Infos) reset() {
	i.Seeders = "0"
	i.Leechers = "0"
	i.Torrentspeed = "0.00MB/s"
}

func defaultInfos(name string) Infos {
	return Infos{
		Name:         name,
		Progress:     "",
		Seeders:      "",
		Leechers:     "",
		Torrentspeed: "",
		ETA:          "",
	}
}
