package pokecache

import "time"

type cacheEntry struct {
	value []byte
	createdAt time.Time
}

type Cache struct {
	cache map[string]cacheEntry
}

func NewCache() Cache {
	return Cache{
		cache: make(map[string]cacheEntry),
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		val,
		time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	cache, ok := c.cache[key]
	return cache.value, ok
}
