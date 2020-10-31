package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/oouxx/proxyaggre/pkg/proxy"
)

var c = cache.New(24*time.Hour, 24*time.Hour)

func GetProxies(key string) proxy.ProxyList {
	result, found := c.Get(key)
	if found {
		return result.(proxy.ProxyList)
	}
	return nil
}

func SetProxies(key string, proxies proxy.ProxyList) {
	c.Set(key, proxies, 0)
}

func SetString(key, value string) {
	c.Set(key, value, 0)
}

func GetString(key string) string {
	result, found := c.Get(key)
	if found {
		return result.(string)
	}
	return ""
}
