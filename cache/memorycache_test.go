package cache

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const loop int = 1000

func TestMemory_GetSet(t *testing.T) {
	n := loop * loop
	m := NewMemoryCache(int64(n), "1s")

	for i := 0; i < n; i++ {
		m.Set(uint64(i), "hello", m.Config.DefaultExpiry)
	}

	start := time.Now()
	_, _ = m.Get(uint64(n / 2))
	end := time.Now()
	assert.LessOrEqual(t, int(end.Sub(start)), 750)
}

func TestMemory_Get(t *testing.T) {
	m := NewMemoryCache(int64(loop), "40µs")
	for i := 0; i < loop; i++ {
		m.Set(uint64(i), fmt.Sprintf("Iteration %d", i), time.Now())
	}

	for i := 0; i < loop; i++ {
		result, _ := m.Get(uint64(i))
		assert.Equal(t, fmt.Sprintf("Iteration %d", i), result)
	}
}

func TestMemory_Set(t *testing.T) {
	m := NewMemoryCache(int64(loop), "40µs")
	for i := 0; i < loop; i++ {
		m.Set(uint64(i), fmt.Sprintf("Iteration %d", i), time.Now())
		m.Set(uint64(i), fmt.Sprintf("Iteration %d", i), time.Now())
		m.Set(uint64(i), fmt.Sprintf("Iteration %d", i), time.Now())
		m.Set(uint64(i), fmt.Sprintf("Iteration %d", i), time.Now())
	}

	for i := 0; i < loop; i++ {
		result, _ := m.Get(uint64(i))
		assert.Equal(t, fmt.Sprintf("Iteration %d", i), result)
	}
}

func TestMemory_Release(t *testing.T) {
	m := NewMemoryCache(int64(loop), "40µs")
	for i := 0; i < loop; i++ {
		m.Set(uint64(i), fmt.Sprintf("Iteration %d", i), time.Now())
	}

	for i := 0; i < loop; i++ {
		m.Release(uint64(i))
		result, _ := m.Get(uint64(i))
		assert.Equal(t, nil, result)
	}
}

func TestMemory_evict(t *testing.T) {
	m := NewMemoryCache(int64(loop), "40µs")
	for i := 0; i < loop*2; i++ {
		m.Set(uint64(i), fmt.Sprintf("Iteration %d", i), time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC))
	}
	assert.NotEqual(t, loop, len(m.store))
}
