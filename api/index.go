package api

import (
	"fmt"
	"github.com/gorilla/mux"
	C "github.com/oouxx/proxyaggre/internal/cache"
	"github.com/oouxx/proxyaggre/internal/cron"
	"github.com/oouxx/proxyaggre/pkg/provider"
	"log"
	"net/http"
)


// GetRouter returns the router for the API
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", Handler).Methods(http.MethodGet)
	return r
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

func respond(w http.ResponseWriter, r *http.Request, body []byte, err error) {
	switch err {
	case nil:
		w.Write(body)
	default:
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}

func marshal(w http.ResponseWriter, r *http.Request, result string) (body []byte, err error){
	w.Header().Set("Content-Type", "text/plain")
	return []byte(result), nil
}

func resolver(r *http.Request) string{
	return "hello"
}
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println(*r)
	body, err := marshal(w, r, resolver(r))
	respond(w, r, body, err)
}