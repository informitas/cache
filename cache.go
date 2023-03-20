package cache

import (
	"errors"
	"sync"
)

type Cache[T any] struct {
	storage map[string]T
	mu *sync.RWMutex
}

// New creates a new cache
func New[T any]() *Cache[T] {
	return &Cache[T]{storage: make(map[string]T), mu: &sync.RWMutex{}}
}

// Set adds an item to the cache. If the key already exists, it overwrites the value
func (c *Cache[T]) Set(key string, value T) {
	c.mu.Lock()
	c.storage[key] = value
	c.mu.Unlock()
}

// Get returns an item from the cache. If the key does not exist, it returns nil
func (c *Cache[T]) Get(key string) T {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.storage[key]
}

// Delete removes an item from the cache. If the key does not exist, it returns an error
func (c *Cache[T]) Delete(key string) error {
	//check if key exists. If not, return an error
	c.mu.Lock()
	if _, ok := c.storage[key]; !ok {
		c.mu.Unlock()
		return errors.New("key does not exist")
	}
	delete(c.storage, key)
	c.mu.Unlock()
	return nil
}

// Clear empties the cache
func (c *Cache[T]) Clear() {
	c.mu.Lock()
	c.storage = make(map[string]T)
	c.mu.Unlock()
}

// Size returns the number of items in the cache
func (c *Cache[T]) Size() int {
	return len(c.storage)
}

// Has checks if a key exists in the cache
func (c *Cache[T]) Has(key string)  bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, ok := c.storage[key]
	return ok
}

// Keys returns a slice of all the keys in the cache
func (c *Cache[T]) Keys() []string {
	keys := make([]string, len(c.storage))
	i := 0
	c.mu.RLock()
	for k := range c.storage {
		keys[i] = k
		i++
	}
	c.mu.RUnlock()
	return keys
}
