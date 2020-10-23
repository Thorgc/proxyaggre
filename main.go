package main

import (
	"flag"
	"fmt"
	"github.com/oouxx/proxyaggre/api"
	"github.com/oouxx/proxyaggre/internal/app"
	"github.com/oouxx/proxyaggre/pkg/proxy"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

var configFilePath = ""

func main() {
	flag.StringVar(&configFilePath, "c", "", "path to config file: config.yaml")
	flag.Parse()

	if configFilePath == "" {
		configFilePath = os.Getenv("CONFIG_FILE")
	}
	if configFilePath == "" {
		configFilePath = "config.yaml"
	}
	err := app.InitConfigAndGetters(configFilePath)
	if err != nil {
		panic(err)
	}

	//database.InitTables()
	proxy.InitGeoIpDB()
	//go app.CrawlGo()
	//go cron.Cron()
	Run()
}

func Run() {
	s := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", "8080"),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  1 * time.Minute,
		Handler:      api.GetRouter(),
	}
	log.Printf("Listening on %s", s.Addr)
	log.Fatal(s.ListenAndServe())
}
