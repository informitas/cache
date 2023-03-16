package cache

import "errors"

type Cache struct {
	storage map[string]interface{}
}

// New creates a new cache
func New() *Cache {
	return &Cache{storage: make(map[string]interface{})}
}

// Set adds an item to the cache. If the key already exists, it overwrites the value
func (c *Cache) Set(key string, value interface{}) {
	c.storage[key] = value
}

// Get returns an item from the cache. If the key does not exist, it returns nil
func (c *Cache) Get(key string) (interface{}, error) {
	if c.Has(key) {
		return c.storage[key], nil
	}
	return nil, errors.New("key does not exist")
}

// Delete removes an item from the cache. If the key does not exist, it returns an error
func (c *Cache) Delete(key string) error {
	//check if key exists. If not, return an error
	if _, ok := c.storage[key]; !ok {
		return errors.New("key does not exist")
	}
	delete(c.storage, key)
	return nil
}

// Clear empties the cache
func (c *Cache) Clear() {
	c.storage = make(map[string]interface{})
}

// Size returns the number of items in the cache
func (c *Cache) Size() int {
	return len(c.storage)
}

// Has checks if a key exists in the cache
func (c *Cache) Has(key string)  bool {
	_, ok := c.storage[key]
	return ok
}
