package database

import (
	"image"
	"sync"
)

type Cache struct {
	data map[string]image.Image
	mu   sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]image.Image),
	}
}

func (c *Cache) Get(key string) (image.Image, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, found := c.data[key]
	return val, found
}

func (c *Cache) Set(key string, value image.Image) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}
