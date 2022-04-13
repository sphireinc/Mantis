package cache

import (
	"github.com/allegro/bigcache"
)

// BigCache primary struct with bigcache pointer and config
type BigCache struct {
	Cache  *bigcache.BigCache
	Config bigcache.Config
}

// Init creates a new Allegro BigCache based on b.Config
func (b *BigCache) Init() error {
	cache, err := bigcache.NewBigCache(b.Config)
	b.Cache = cache
	return err
}
