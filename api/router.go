package api


import (
	"fmt"
	"github.com/gorilla/mux"
	C "github.com/oouxx/proxyaggre/internal/cache"
	"github.com/oouxx/proxyaggre/internal/cron"
	_ "github.com/oouxx/proxyaggre/internal/cron"
	"github.com/oouxx/proxyaggre/pkg/provider"
	"log"
	"net/http"
	"time"
)


// GetRouter returns the router for the API
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods(http.MethodGet)
	r.HandleFunc("/vmess/sub", vmessSub).Methods(http.MethodGet)
	r.HandleFunc("/ss/sub", ssSub).Methods(http.MethodGet)
	r.HandleFunc("/ssr/sub", ssrSub).Methods(http.MethodGet)
	r.HandleFunc("/sip002/sub", sip002ub).Methods(http.MethodGet)
	r.HandleFunc("/cron", runCron).Methods(http.MethodGet)
	return r
}
func Run() {
	s := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", "8080"),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  1 * time.Minute,
		Handler:      GetRouter(),
	}
	log.Printf("Listening on %s", s.Addr)
	log.Fatal(s.ListenAndServe())
}

func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>哈喽啊!首页待完善</h1>")
}

func vmessSub(w http.ResponseWriter, r *http.Request){
	proxies := C.GetProxies("proxies")
	vmessSub := provider.VmessSub{
		provider.Base{
			Proxies: &proxies,
			Types:   "vmess",
		},
	}
	fmt.Fprintf(w, vmessSub.Provide())
}
func ssSub(w http.ResponseWriter, r *http.Request){
	proxies := C.GetProxies("proxies")
	ssSub := provider.SSSub{
		provider.Base{
			Proxies: &proxies,
			Types:   "ss",
		},
	}
	fmt.Fprintf(w, ssSub.Provide())
}

func ssrSub(w http.ResponseWriter, r *http.Request){
	proxies := C.GetProxies("proxies")
	ssrSub := provider.SSRSub{
		provider.Base{
			Proxies: &proxies,
			Types: "ssr",
		},
	}
	fmt.Fprint(w, ssrSub.Provide())
}
func sip002ub(w http.ResponseWriter, r *http.Request){
	proxies := C.GetProxies("proxies")
	sip002Sub := provider.SIP002Sub{
		provider.Base{
			Proxies: &proxies,
			Types: "ss",
		},
	}
	fmt.Fprint(w, sip002Sub.Provide())
}

func runCron(w http.ResponseWriter, r *http.Request){
	cron.CrawlTask()
	fmt.Fprintf(w, "<h1>正在运行cron任务</h1>")
}