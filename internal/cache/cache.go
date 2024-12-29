package cache

import (
	"time"
	"sync"
)

type Cache struct {
	cache map[string]CacheEntry
	rlock sync.RWMutex
	reapInterval time.Duration
}

type CacheEntry struct {
	createdAt time.Time
	val []byte
}

func (c *Cache) Add(key string, val []byte) {

}

func (c *Cache) Get(key string) ([]byte, bool) {
	return nil, false
}

func (c *Cache) ReapLoop() {
	
}

func NewCache(reapInterval time.Duration) Cache {
	return Cache{
		reapInterval: reapInterval,
	}
}