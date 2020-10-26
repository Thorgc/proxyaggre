package api

import (
	"fmt"
	"github.com/oouxx/proxyaggre/internal/cron"
	"net/http"
)

func CronHandler(w http.ResponseWriter, r *http.Request){
	cron.CrawlTask()
	fmt.Fprintf(w, "Cron job is running.")
}
