package cache

import (
	"errors"
	"time"

	"github.com/bluehawk27/assignment/config"
	"github.com/karlseguin/ccache"
)

// Cache : represents in Memory Cache
type Cache struct {
	Cache  *ccache.Cache
	Expiry int64
}

// NewCache : Instantiate In Memory Cache with a size and Expiry
func NewCache() *Cache {
	conf := config.GetCCacheConfig()
	ccacheConfig := ccache.Configure().MaxSize(conf.Capacity)
	newCache := ccache.New(ccacheConfig)
	c := Cache{
		Cache:  newCache,
		Expiry: conf.Expiry,
	}
	return &c
}

// Set : Set Key - Value in memory
func (c *Cache) Set(key string, value interface{}) {
	ttl := time.Duration(c.Expiry) * time.Second
	c.Cache.Set(key, value, ttl)
}

// Get : Get Value from Key from in memory structure
func (c *Cache) Get(key string) (interface{}, error) {
	val := c.Cache.Get(key)
	if val == nil {
		return nil, errors.New("Item not set in Cache")
	}
	resp := val.Value()
	return resp, nil
}
