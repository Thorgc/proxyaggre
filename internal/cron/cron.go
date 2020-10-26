package cron

import (
	"github.com/oouxx/proxyaggre/internal/app"
	"log"
	"os"
	"runtime"
)

//func Cron() {
//	_ = gocron.Every(15).Minutes().Do(crawlTask)
//	<-gocron.Start()
//}

func CrawlTask() {
	var configFilePath = os.Getenv("CONFIG_FILE")
	log.Println(configFilePath)
	_ = app.InitConfigAndGetters(configFilePath)
	app.CrawlGo()
	app.Getters = nil
	runtime.GC()
}
