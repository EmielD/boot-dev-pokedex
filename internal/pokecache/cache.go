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
	mu           sync.Mutex
	cacheEntries map[string]cacheEntry
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cacheEntries: make(map[string]cacheEntry),
		mu:           sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cacheEntries[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cacheEntry, ok := c.cacheEntries[key]
	if !ok {
		return nil, false
	}

	return cacheEntry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.mu.Lock()
		for key, cacheEntry := range c.cacheEntries {
			timeSinceCreation := time.Since(cacheEntry.createdAt)
			if timeSinceCreation > interval {
				delete(c.cacheEntries, key)
			}
		}
		c.mu.Unlock()
	}
}
