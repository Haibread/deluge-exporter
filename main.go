package main

import (
	"net/http"
	"os"

	"github.com/Haibread/deluge-exporter/collector"
	"github.com/Haibread/deluge-exporter/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)

	http.HandleFunc("/metrics", metricsHandler)
	http.ListenAndServe(":2112", nil)

}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	// Read config file
	logrus.Debug("Reading config file")
	configfile, err := os.ReadFile("config.yml")
	if err != nil {
		logrus.Error(err)
	}
	// Parse config file
	logrus.Debug("Parsing config file")
	var config config.Config
	err = yaml.Unmarshal(configfile, &config)
	if err != nil {
		logrus.Error(err)
	}

	registry := prometheus.NewRegistry()
	registry.MustRegister(collectors.NewGoCollector())
	// Register collectors
	logrus.Debug("Registering collectors")
	for _, client := range config.DelugeClients {
		logrus.Debugf("%+v\n", client)
		collector := collector.NewDelugeCollector(client)
		prometheus.WrapRegistererWith(prometheus.Labels{"instance_name": client.Name, "instance": client.Host}, registry).MustRegister(collector)
	}

	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)
}
