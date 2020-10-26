package api

import (
	"github.com/gorilla/mux"
	C "github.com/oouxx/proxyaggre/internal/cache"
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

func vmessSub() (subString string){
	proxies := C.GetProxies("proxies")
	vmessSub := provider.VmessSub{
		provider.Base{
			Proxies: &proxies,
			Types:   "vmess",
		},
	}
	return vmessSub.Provide()
}
func ssSub() (subString string){
	proxies := C.GetProxies("proxies")
	ssSub := provider.SSSub{
		provider.Base{
			Proxies: &proxies,
			Types:   "ss",
		},
	}
	return ssSub.Provide()
}

func ssrSub() (subString string){
	proxies := C.GetProxies("proxies")
	ssrSub := provider.SSRSub{
		provider.Base{
			Proxies: &proxies,
			Types: "ssr",
		},
	}
	return ssrSub.Provide()
}
func sip002ub() (subString string){
	proxies := C.GetProxies("proxies")
	sip002Sub := provider.SIP002Sub{
		provider.Base{
			Proxies: &proxies,
			Types: "ss",
		},
	}
	return sip002Sub.Provide()
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
	return vmessSub()
}
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println(*r)
	body, err := marshal(w, r, resolver(r))
	respond(w, r, body, err)
}