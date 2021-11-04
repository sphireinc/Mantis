package cache

import (
	"github.com/allegro/bigcache"
)

type BigCache struct {
	Cache  *bigcache.BigCache
	Config bigcache.Config
}

func (b *BigCache) Init() error {
	cache, err := bigcache.NewBigCache(b.Config)
	b.Cache = cache
	return err
}
