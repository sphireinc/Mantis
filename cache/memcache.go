package cache

import (
	mantisError "github.com/sphireinc/mantis/error"
	"github.com/victorspringer/http-cache"
	"github.com/victorspringer/http-cache/adapter/memory"
	"time"
)

type memCache struct {
	Client       *cache.Client
	Algorithm    memory.Algorithm
	Capacity     int
	RefreshKey   string
	memCacheTime time.Duration
}

// NewMemCache creates a new MemCache instance
func NewMemCache(algorithm memory.Algorithm, capacity int, refreshKey string, cacheTime time.Duration) *memCache {
	return &memCache{
		Algorithm:    algorithm,
		Capacity:     capacity,
		RefreshKey:   refreshKey,
		memCacheTime: cacheTime * time.Minute,
	}
}

// StartCache starts our in-memory cache.
func (m *memCache) Init() error {
	memoryCache, err := memory.NewAdapter(
		memory.AdapterWithAlgorithm(m.Algorithm),
		memory.AdapterWithCapacity(m.Capacity),
	)
	if err != nil {
		mantisError.HandleError("Error starting memory cache", err)
		return err
	}

	m.Client, err = cache.NewClient(
		cache.ClientWithAdapter(memoryCache),
		cache.ClientWithTTL(m.memCacheTime),
		cache.ClientWithRefreshKey(m.RefreshKey),
	)

	if err != nil {
		mantisError.HandleError("Error starting http cache", err)
		return err
	}

	return nil
}
