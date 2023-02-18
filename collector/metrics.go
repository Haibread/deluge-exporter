package collector

import "github.com/prometheus/client_golang/prometheus"

var daemonversion = prometheus.NewDesc("deluge_daemon_version", "Deluge daemon version", []string{"version"}, nil)

var torrents_numbers = prometheus.NewDesc("deluge_torrents_numbers", "Number of torrents", nil, nil)
