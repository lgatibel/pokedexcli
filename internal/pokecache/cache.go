package pokecache

import (
	"fmt"
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
	go cache.reapLoop()
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

func (c *Cache) Get(key string) ([]byte, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if data, ok := c.Datas[key]; ok {
		return data.val, nil
	}
	return nil, fmt.Errorf("no cache entry for the key: %s", key)
}

func (c *Cache) reapLoop() {
	defer c.mu.Unlock()
	c.mu.Lock()

	for key, entry := range c.Datas {
		if time.Since(entry.createdAt) > c.Interval {
			delete(c.Datas, key)
			fmt.Println(c.Datas)
		}
	}
}
