package cache

import (
	"sync"
	"time"
)

type item struct {
	value        any
	lastAccessed time.Time
	expiration   time.Time
}

type Memory struct {
	mutex    sync.RWMutex
	capacity int64
	store    map[uint64]item
	Config   struct {
		Expiry string
	}
}

func NewMemoryCache(capacity int64, expiry string) *Memory {
	return &Memory{
		capacity: capacity,
		store:    make(map[uint64]item),
		Config:   struct{ Expiry string }{Expiry: expiry},
	}
}

func (m *Memory) Get(key uint64) (any, bool) {
	m.mutex.RLock()
	item, ok := m.store[key]
	m.mutex.RUnlock()

	if ok {
		m.Set(key, item.value, item.expiration)
		return item.value, true
	}
	return nil, false
}

func (m *Memory) Set(key uint64, value any, expiration time.Time) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	toStore := item{
		value:        value,
		lastAccessed: time.Now(),
		expiration:   expiration,
	}

	// check for existence of the key, overwrite new value and return
	if _, ok := m.store[key]; ok {
		m.store[key] = toStore
		return
	}

	// Make sure we have capacity, if not then evict
	if int64(len(m.store)) >= m.capacity {
		m.evict()
	}

	m.store[key] = toStore
}

func (m *Memory) Release(key uint64) {
	m.mutex.RLock()
	_, ok := m.store[key]
	m.mutex.RUnlock()
	if ok {
		m.mutex.Lock()
		delete(m.store, key)
		m.mutex.Unlock()
	}
}

// evict records from memory on an LRU basis
func (m *Memory) evict() {
	now := time.Now()
	accessCutoff, _ := time.ParseDuration(m.Config.Expiry)

	for k, v := range m.store {
		if now.Sub(v.lastAccessed) > accessCutoff {
			delete(m.store, k)
		}
	}
}
