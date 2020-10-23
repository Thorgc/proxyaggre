package api

import (
	"fmt"
	"github.com/gorilla/mux"
	C "github.com/oouxx/proxyaggre/internal/cache"
	"github.com/oouxx/proxyaggre/internal/cron"
	"github.com/oouxx/proxyaggre/pkg/provider"
	"net/http"
)


// GetRouter returns the router for the API
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", Index).Methods(http.MethodGet)
	r.HandleFunc("/vmess/sub", VmessSub).Methods(http.MethodGet)
	r.HandleFunc("/ss/sub", SsSub).Methods(http.MethodGet)
	r.HandleFunc("/ssr/sub", SsrSub).Methods(http.MethodGet)
	r.HandleFunc("/sip002/sub", Sip002ub).Methods(http.MethodGet)
	r.HandleFunc("/cron", RunCron).Methods(http.MethodGet)
	return r
}

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>哈喽啊!首页待完善</h1>")
}

func VmessSub(w http.ResponseWriter, r *http.Request){
	proxies := C.GetProxies("proxies")
	vmessSub := provider.VmessSub{
		provider.Base{
			Proxies: &proxies,
			Types:   "vmess",
		},
	}
	fmt.Fprintf(w, vmessSub.Provide())
}
func SsSub(w http.ResponseWriter, r *http.Request){
	proxies := C.GetProxies("proxies")
	ssSub := provider.SSSub{
		provider.Base{
			Proxies: &proxies,
			Types:   "ss",
		},
	}
	fmt.Fprintf(w, ssSub.Provide())
}

func SsrSub(w http.ResponseWriter, r *http.Request){
	proxies := C.GetProxies("proxies")
	ssrSub := provider.SSRSub{
		provider.Base{
			Proxies: &proxies,
			Types: "ssr",
		},
	}
	fmt.Fprint(w, ssrSub.Provide())
}
func Sip002ub(w http.ResponseWriter, r *http.Request){
	proxies := C.GetProxies("proxies")
	sip002Sub := provider.SIP002Sub{
		provider.Base{
			Proxies: &proxies,
			Types: "ss",
		},
	}
	fmt.Fprint(w, sip002Sub.Provide())
}

func RunCron(w http.ResponseWriter, r *http.Request){
	cron.CrawlTask()
	fmt.Fprintf(w, "<h1>正在运行cron任务</h1>")
}