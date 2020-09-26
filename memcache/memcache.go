package memcache

import (
	mantisError "github.com/sphireco/mantis/error"
	"github.com/victorspringer/http-cache"
	"github.com/victorspringer/http-cache/adapter/memory"
	"time"
)

type MemCache struct {
	Client     *cache.Client
	Algorithm  memory.Algorithm
	Capacity   int
	RefreshKey string
}

// StartCache starts our in-memory LRU cache.
func (m *MemCache) Init(MemCacheTime time.Duration) error {
	memoryCache, err := memory.NewAdapter(
		memory.AdapterWithAlgorithm(memory.LFU),
		memory.AdapterWithCapacity(10000000),
	)
	if err != nil {
		mantisError.HandleError("Error starting memory cache", err)
		return err
	}

	m.Client, err = cache.NewClient(
		cache.ClientWithAdapter(memoryCache),
		cache.ClientWithTTL(MemCacheTime*time.Minute),
		cache.ClientWithRefreshKey("opn"),
	)

	if err != nil {
		mantisError.HandleError("Error starting http cache", err)
		return err
	}

	return nil
}
