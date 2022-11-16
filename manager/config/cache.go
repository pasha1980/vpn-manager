package config

import "vpn-manager/services/cache"

var Cache *cache.Cache

func initCache() {
	data := make(map[string]interface{})
	Cache = &cache.Cache{
		Temp: data,
		Data: data,
	}
}
