package cache

import (
	"github.com/allegro/bigcache"
	mantisError "github.com/sphireco/mantis/error"
)

type bigCache struct {
	Cache  *bigcache.BigCache
	Config bigcache.Config
}

// NewBigCache
func NewBigCache(config bigcache.Config) *bigCache {
	return &bigCache{
		Config: config,
	}
}

// Init
func (b *bigCache) Init() error {
	var err error
	b.Cache, err = bigcache.NewBigCache(b.Config)

	if err != nil {
		mantisError.HandleFatalError(err)
		return err
	}

	return nil
}
