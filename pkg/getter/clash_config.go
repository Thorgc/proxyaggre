package getter

import (
	"io/ioutil"
	"sync"

	"github.com/oouxx/proxyaggre/pkg/proxy"
	"github.com/oouxx/proxyaggre/pkg/tool"
)

func init() {
	Register("clash_config", NewClashConfigGetter)
}

type ClashConfigGetter struct {
	Url string
}

func (c *ClashConfigGetter) Get() proxy.ProxyList {
	resp, err := tool.GetHttpClient().Get(c.Url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	return ParseClashProxyFromYamlConfig(body)
}

func (c *ClashConfigGetter) Get2Chan(pc chan proxy.Proxy, wg *sync.WaitGroup) {
	panic("implement me")
}

//func (s *Subscribe) Get() proxy.ProxyList {
//	resp, err := tool.GetHttpClient().Get(s.Url)
//	if err != nil {
//		return nil
//	}
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return nil
//	}
//
//	nodesString, err := tool.Base64DecodeString(string(body))
//	if err != nil {
//		return nil
//	}
//	nodesString = strings.ReplaceAll(nodesString, "\t", "")
//
//	nodes := strings.Split(nodesString, "\n")
//	return StringArray2ProxyArray(nodes)
//}
//
//func (s *Subscribe) Get2Chan(pc chan proxy.Proxy, wg *sync.WaitGroup) {
//	defer wg.Done()
//	nodes := s.Get()
//	log.Printf("STATISTIC: Subscribe\tcount=%d\turl=%s\n", len(nodes), s.Url)
//	for _, node := range nodes {
//		pc <- node
//	}
//}
//
func NewClashConfigGetter(options tool.Options) (getter Getter, err error) {
	urlInterface, found := options["url"]
	if found {
		url, err := AssertTypeStringNotNull(urlInterface)
		if err != nil {
			return nil, err
		}
		return &ClashConfigGetter{
			Url: url,
		}, nil
	}
	return nil, ErrorUrlNotFound
}
