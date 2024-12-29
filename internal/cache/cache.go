package cache

import (
	"sync"
	"time"
)

type Cache struct {
	cache        map[string]CacheEntry
	rlock        sync.RWMutex
	reapInterval time.Duration
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.rlock.Lock()
	defer c.rlock.Unlock()
	c.cache[key] = CacheEntry{
		time.Now(),
		val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.rlock.RLock()
	defer c.rlock.RUnlock()
	val, ok := c.cache[key]
	if !ok {
		return nil, ok
	}
	return val.val, ok
}

func (c *Cache) ReapLoop() {
	ticker := time.NewTicker(c.reapInterval)
	defer ticker.Stop()
	for {
		c.reapNow(<-ticker.C)
	}

}

func (c *Cache) reapNow(t time.Time) {
	c.rlock.Lock()
	defer c.rlock.Unlock()

	cutoffTime := t.Add(c.reapInterval * -1)

	for k, v := range c.cache {
		if v.createdAt.Before(cutoffTime) {
			delete(c.cache, k)
		}
	}

}

func NewCache(reapInterval time.Duration) Cache {
	cache := Cache{
		cache:        make(map[string]CacheEntry),
		rlock:        sync.RWMutex{},
		reapInterval: reapInterval,
	}
	go cache.ReapLoop()
	return cache
}
