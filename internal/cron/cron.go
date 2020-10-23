package cron

import (
	"runtime"

	"github.com/oouxx/proxyaggre/internal/app"
)

//func Cron() {
//	_ = gocron.Every(15).Minutes().Do(crawlTask)
//	<-gocron.Start()
//}

func CrawlTask() {
	_ = app.InitConfigAndGetters("")
	app.CrawlGo()
	app.Getters = nil
	runtime.GC()
}
