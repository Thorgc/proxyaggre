package cron

import (
	"github.com/oouxx/proxyaggre/internal/app"
	"github.com/oouxx/proxyaggre/pkg/proxy"
	"os"
	"runtime"
)

//func Cron() {
//	_ = gocron.Every(15).Minutes().Do(crawlTask)
//	<-gocron.Start()
//}

func CrawlTask() {
	var configFilePath = os.Getenv("CONFIG_FILE")
	// 初始化geoip db
	proxy.InitGeoIpDB()
	_ = app.InitConfigAndGetters(configFilePath)
	app.CrawlGo()
	app.Getters = nil
	runtime.GC()
}
