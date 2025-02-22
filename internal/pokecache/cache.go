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

func NewCache(interval int) Cache {
	nCache := Cache {
		cache: map[string]cacheEntry{},
		mu:    sync.Mutex{},
	}
	nCache.reapLoop(interval)
	return nCache
}

func (c Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry {
		createdAt: time.Now(),
		val: val,
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.cache[key]
	if ok {
		return entry.val, true
	}
	return nil, false
}

func (c Cache) reapLoop(interval int) {
	t := time.Now()
	reapTime := t.Add(time.Duration(-interval) * time.Minute)

	for {
		c.mu.Lock()
		for k := range c.cache {
			if !c.cache[k].createdAt.Before(reapTime) {
				delete(c.cache, k)
			}
		}
		c.mu.Unlock()
		reapFreq := 5
		time.Sleep(time.Duration(reapFreq) * time.Second)
	}
}
