package pokeapi

import (
	"sync"
	"time"
)

type Cache struct {
	cacheE       map[string]cacheEntry
	protMapMutex sync.Mutex
	interval     time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	newC := &Cache{
		cacheE:   make(map[string]cacheEntry),
		interval: interval,
	}
	newC.reapLoop()
	return newC
}

func (c *Cache) Add(key string, val []byte) {
	c.protMapMutex.Lock()
	defer c.protMapMutex.Unlock()
	c.cacheE[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

}
func (c *Cache) Get(key string) ([]byte, bool) {
	// The bool should be true if the entry was found and false if it wasn't.
	c.protMapMutex.Lock()
	defer c.protMapMutex.Unlock()
	entry, found := c.cacheE[key]
	if !found {
		return nil, false
	}
	return entry.val, true
}
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		c.protMapMutex.Lock()
		now := time.Now()
		for k, v := range c.cacheE {
			if now.Sub(v.createdAt) > c.interval {
				delete(c.cacheE, k)
			}
		}
		defer c.protMapMutex.Unlock()
	}

}
