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
		Expiry        string
		DefaultExpiry time.Time
	}
}

func NewMemoryCache(capacity int64, expiry string) *Memory {
	m := &Memory{
		capacity: capacity,
		store:    make(map[uint64]item),
		Config: struct {
			Expiry        string
			DefaultExpiry time.Time
		}{
			Expiry:        expiry,
			DefaultExpiry: time.Now().Add(time.Duration(100)),
		},
	}

	if m.Config.Expiry == "" {
		m.Config.Expiry = m.Config.DefaultExpiry.String()
	}
	return m
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
	_, ok := m.store[key]
	m.mutex.Unlock()

	toStore := item{
		value:        value,
		lastAccessed: time.Now(),
		expiration:   expiration,
	}

	// check for existence of the key, overwrite new value and return
	if ok {
		expDiff := m.store[key].lastAccessed.Sub(m.store[key].expiration)
		if exp, _ := time.ParseDuration(m.Config.Expiry); exp > expDiff {
			m.Release(key)
		}
		m.store[key] = toStore
		return
	}

	// Make sure we have capacity, if not then evict
	if int64(len(m.store)) >= m.capacity {
		m.evict()
	}

	m.mutex.Lock()
	m.store[key] = toStore
	m.mutex.Unlock()
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
