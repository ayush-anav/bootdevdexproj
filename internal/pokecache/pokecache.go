package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte // data we caching
}

type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}
	// everytime a cache is created, we will call ReapLoop
	// ReapLoop ensures that TTL is adhered to (removes old cache)
	go c.ReapLoop(interval)
	return c
}

// adding to cache struct
func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
}

// get from cache
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	hit, ok := c.cache[key]
	if !ok {
		return []byte{}, ok
	}
	return hit.val, ok
}

func (c *Cache) ReapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cache, k)
		}
	}
}
