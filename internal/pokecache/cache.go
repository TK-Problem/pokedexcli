package pokecache

import (
	"sync"
	"time"
)

// Cache is a concurrency-safe map of raw response bodies keyed by request URL.
type Cache struct {
	mu      sync.Mutex
	entries map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// NewCache returns a cache whose entries are reaped once they grow older than
// interval. The reaping goroutine runs for the life of the process.
func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]cacheEntry),
	}

	go c.reapLoop(interval)

	return c
}

// Add stores val under key, replacing any existing entry.
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

// Get returns the value stored under key. The bool reports whether it was found.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	return entry.val, ok
}

// reapLoop evicts stale entries once per interval, forever.
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.reap(interval)
	}
}

// reap deletes every entry created more than interval ago.
func (c *Cache) reap(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cutoff := time.Now().Add(-interval)
	for key, entry := range c.entries {
		if entry.createdAt.Before(cutoff) {
			delete(c.entries, key)
		}
	}
}
