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
	Expiry        time.Duration
	DefaultExpiry time.Duration
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
	expiryDuration, err := time.ParseDuration(expiry)
	if err != nil {
		expiryDuration = time.Second
	}
	m := &Memory{
		capacity: capacity,
		store:    make(map[uint64]item),
		Config: memoryConfig{
			Expiry:        expiryDuration,
			DefaultExpiry: 100 * time.Millisecond,
		},
	}
	return m
}

// Get the value associated with a key in our cache
func (m *Memory) Get(key uint64) (any, bool) {
	m.mutex.RLock()
	item, ok := m.store[key]
	m.mutex.RUnlock()

	if ok {
		m.checkExpireAndUpdate(key, item)
		return item.value, true
	}

	return nil, false
}

// checkExpireAndUpdate checks if KV has expired, if not updates toStore item
// returns false if released, true if updated
func (m *Memory) checkExpireAndUpdate(key uint64, toStore item) int {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	existingItem, ok := m.store[key]
	status := NOOP

	if !ok {
		m.store[key] = toStore
		status = CREATED
	} else {
		// Correctly check against the expiration time
		if time.Now().After(existingItem.expiration) {
			m.Release(key)
			status = RELEASED
		} else {
			toStore.lastAccessed = time.Now()
			m.store[key] = toStore
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
	delete(m.store, key)
	return
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
	m.mutex.Lock()
	defer m.mutex.Unlock()

	now := time.Now()
	for k, v := range m.store {
		if now.Sub(v.lastAccessed) > m.Config.Expiry {
			delete(m.store, k)
		}
	}
}
