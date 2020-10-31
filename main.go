package main

import (
	"flag"
	"fmt"
	_ "net/http/pprof"
	"os"

	"github.com/oouxx/proxyaggre/api"
	"github.com/oouxx/proxyaggre/internal/app"
	"github.com/oouxx/proxyaggre/internal/cron"
	"github.com/oouxx/proxyaggre/internal/database"
	"github.com/oouxx/proxyaggre/pkg/proxy"
)

var configFilePath = ""

func main() {
	// support -c config file path
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

	database.InitTables()
	proxy.InitGeoIpDB()
	fmt.Println("Do the first crawl...")
	go app.CrawlGo()
	go cron.Cron()
	api.Run()
}
