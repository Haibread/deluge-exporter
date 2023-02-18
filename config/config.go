package config

type Config struct {
	DelugeClients []DelugeClient `yaml:"deluge_clients"`
}

type DelugeClient struct {
	Name              string `yaml:"name"`
	Host              string `yaml:"host"`
	Port              int    `yaml:"port"`
	Username          string `yaml:"username"`
	Password          string `yaml:"password"`
	PerTorrentMetrics bool   `yaml:"per_torrent_metrics"`
}
