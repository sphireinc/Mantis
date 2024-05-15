package cache

import (
	"sync"
	"time"
)

// RELEASED is a status enum for a released item
// UPDATED is a status enum for an updated value in an existing kv item
// CREATED is a status enum for a created kv item
// NOOP is a status for no operation
const (
	RELEASED = iota
	UPDATED
	CREATED
	NOOP
)

// item is an item and its metadata in our cache
type item struct {
	value        any
	lastAccessed time.Time
	expiration   time.Time
}

// memoryConfig holds our expiry configuration
type memoryConfig struct {
	Expiry        string
	DefaultExpiry time.Time
}

// Memory holds our cache
type Memory struct {
	mutex    sync.RWMutex
	capacity int64
	store    map[uint64]item
	Config   memoryConfig
}

// NewMemoryCache creates and returns a new in memory cache
func NewMemoryCache(capacity int64, expiry string) *Memory {
	m := &Memory{
		capacity: capacity,
		store:    make(map[uint64]item),
		Config: memoryConfig{
			Expiry:        expiry,
			DefaultExpiry: time.Now().Add(time.Duration(100)),
		},
	}

	if m.Config.Expiry == "" {
		m.Config.Expiry = m.Config.DefaultExpiry.String()
	}
	return m
}

// Get the value associated with a key in our cache
func (m *Memory) Get(key uint64) (any, bool) {
	m.mutex.Lock()
	_, ok := m.store[key]
	m.mutex.Unlock()

	if ok {
		m.checkExpireAndUpdate(key, m.store[key])
		return m.store[key].value, true
	}

	return nil, false
}

// checkExpireAndUpdate checks if KV has expired, if not updates toStore item
// returns false if released, true if updated
func (m *Memory) checkExpireAndUpdate(key uint64, toStore item) int {
	m.mutex.Lock()
	_, ok := m.store[key]
	status := NOOP
	m.mutex.Unlock()

	if !ok {
		m.store[key] = toStore
		status = CREATED
	} else {
		// check for existence of the key, overwrite new value and return
		expirationDiff := m.store[key].lastAccessed.Sub(m.store[key].expiration)
		expiry, _ := time.ParseDuration(m.Config.Expiry)
		if expirationDiff > expiry {
			m.Release(key)
			status = RELEASED
		} else {
			m.mutex.Lock()
			toStore.lastAccessed = time.Now()
			m.store[key] = toStore
			m.mutex.Unlock()
			status = UPDATED
		}
	}

	return status
}

// Set a new kv pair in our memory cache
func (m *Memory) Set(key uint64, value any, expiration time.Time) {
	m.triggerEvict()
	m.checkExpireAndUpdate(key, item{
		value:        value,
		lastAccessed: time.Now(),
		expiration:   expiration,
	})
}

// Release an item from our memory cache
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

// triggerEvict checks if we need to evict data, and commences eviction if so
func (m *Memory) triggerEvict() {
	// Make sure we have capacity, if not then evict
	if int64(len(m.store)) >= m.capacity {
		m.evict()
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
