package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	data map[string]cacheEntry
	mu   *sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	ticker := time.NewTicker(interval)

	cache := &Cache{
		data: map[string]cacheEntry{},
		mu:   &sync.Mutex{},
	}

	go func() {
		for range ticker.C {
			cache.mu.Lock()
			for k, v := range cache.data {
				if time.Since(v.createdAt) >= interval {
					delete(cache.data, k)
				}
			}
			cache.mu.Unlock()
		}
	}()

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	ce := cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
	c.data[key] = ce
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	dataByKey, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return dataByKey.val, true
}

func (c *Cache) Remove(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, key)
}
