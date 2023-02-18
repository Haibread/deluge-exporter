package collector

import (
	"fmt"

	"github.com/Haibread/deluge-exporter/config"
	delugeclient "github.com/gdm85/go-libdeluge"
	"github.com/prometheus/client_golang/prometheus"
)

type DelugeCollector struct {
	client *delugeclient.ClientV2
	config config.DelugeClient
}

func NewDelugeCollector(c config.DelugeClient) *DelugeCollector {

	client := delugeclient.NewV2(delugeclient.Settings{
		Hostname: c.Host,
		Port:     uint(c.Port),
		Login:    c.Username,
		Password: c.Password,
	})
	err := client.Connect()
	if err != nil {
		fmt.Println(err)
	}

	d := &DelugeCollector{
		client: client,
		config: c,
	}
	return d
}

func (e *DelugeCollector) Collect(ch chan<- prometheus.Metric) {
	//ch <- prometheus.MustNewConstMetric(torrentsDesc, prometheus.GaugeValue, 1)
	e.collectDaemonVersion(ch)
	e.collectTorrentsDetails(ch)
}

func (e *DelugeCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- daemonversion
	ch <- torrents_numbers
}

func (e *DelugeCollector) collectDaemonVersion(ch chan<- prometheus.Metric) {
	daemonVersion, err := e.client.DaemonVersion()
	if err != nil {
		fmt.Println(err)
	}
	ch <- prometheus.MustNewConstMetric(daemonversion, prometheus.GaugeValue, float64(1), daemonVersion)
}

func (e *DelugeCollector) collectTorrentsDetails(ch chan<- prometheus.Metric) {
	torrents, err := e.client.TorrentsStatus(delugeclient.StateUnspecified, nil)
	if err != nil {
		fmt.Println(err)
	}
	ch <- prometheus.MustNewConstMetric(torrents_numbers, prometheus.GaugeValue, float64(len(torrents)))
}
