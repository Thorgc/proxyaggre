package main

import (
	"flag"
	"fmt"
	"github.com/oouxx/proxyaggre/api"
	"github.com/oouxx/proxyaggre/internal/app"
	"github.com/oouxx/proxyaggre/pkg/proxy"
	"net/http"
	_ "net/http/pprof"
	"os"
)

var configFilePath = ""

func main() {
	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()

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
	fmt.Println("Do the first crawl...")
	go app.CrawlGo()
	//go cron.Cron()
	api.Run()
}
