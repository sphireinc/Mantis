package cache

import (
	"github.com/allegro/bigcache"
	mantisError "github.com/sphireco/mantis/error"
)

type BigCache struct {
	Cache *bigcache.BigCache
	Config bigcache.Config
}

func (b *BigCache) Init() (*bigcache.BigCache, error) {
	cache, err := bigcache.NewBigCache(b.Config)

	if err != nil {
		mantisError.HandleFatalError(err)
		return nil, err
	}

	return cache, nil
}
