package pokecache

import (
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt   time.Time
	val         []byte
}

type Cache struct {
	cache      map[string]cacheEntry
	mu         sync.Mutex
}

func NewCache(interval int) {
	nCache := cache {
		cache: cacheEntry{},
		mu:    &sync.Mutex{},
	}
	nCache.reapLoop()
}

func (c cache) Add(key string, val []byte) {
	c.cache[key] = val
	return nil
}

func (c cache) Get(key string) ([]byte, bool) {
	val, ok := c.cache[key]
	if ok {
		return val, true
	}
	return nil, false
}

func (c cache) reapLoop(interval int) {
	t := time.Now()
	reapTime := t.Add(time.Duration(-interval) * time.Minute)

	for {
		c.mu.Lock()
		for k := range c.cache {
			if c.cache[k].createdAt >= reapTime {
				delete(c.cache, k)
			}
		}
		c.mu.Unlock()
		reapFreq := 5
		time.Sleep(time.Duration(reapFreq) * time.Second)
	}
}
