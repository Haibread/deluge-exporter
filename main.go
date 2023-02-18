package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Haibread/deluge-exporter/collector"
	"github.com/Haibread/deluge-exporter/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/yaml.v2"
)

func main() {

	http.HandleFunc("/metrics", metricsHandler)
	http.ListenAndServe(":2112", nil)

}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	// Read config file
	fmt.Println("Reading config file")
	configfile, err := os.ReadFile("config.yml")
	if err != nil {
		fmt.Println(err)
	}
	// Parse config file
	fmt.Println("Parsing config file")
	var config config.Config
	err = yaml.Unmarshal(configfile, &config)
	if err != nil {
		fmt.Println(err)
	}

	registry := prometheus.NewPedanticRegistry()
	registry.MustRegister(collectors.NewGoCollector())
	// Register collectors
	fmt.Println("Registering collectors")
	for _, client := range config.DelugeClients {
		fmt.Printf("%+v\n", client)
		collector := collector.NewDelugeCollector(client)
		prometheus.WrapRegistererWith(prometheus.Labels{"instance_name": client.Name}, registry).MustRegister(collector)
	}

	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)
}
