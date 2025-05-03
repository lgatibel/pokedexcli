package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Icache interface {
	reapLoop(Cache)
}

type Cache struct {
	Interval time.Duration
	Datas    map[string]cacheEntry
	mu       sync.RWMutex
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		Interval: interval,
		Datas:    make(map[string]cacheEntry),
	}
	cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, data []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.Datas == nil {
		c.Datas = make(map[string]cacheEntry)
	}
	c.Datas[key] = cacheEntry{
		createdAt: time.Now(),
		val:       data,
	}
	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if data, ok := c.Datas[key]; ok {
		return data.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop() {
	test := make(chan time.Time)
	go func() {
		tick := time.Tick(c.Interval)
		for t := range tick {
			test <- t
		}
	}()

	go func() {
		for range test {
			for key, entry := range c.Datas {
				c.mu.Lock()
				if time.Since(entry.createdAt) >= c.Interval {
					delete(c.Datas, key)
				}
				c.mu.Unlock()
			}
		}
	}()
}
