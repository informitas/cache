package internal

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type CacheStore[T any] struct {
	Data        T
	Expiration  time.Time
	Immutable		bool
}

type Cache[T any] struct {
	storage map[string]CacheStore[T]
	mu *sync.RWMutex
	TTLTicker *time.Ticker
	showExpiredLog bool
	logger io.Writer
}

type CacheOptions struct {
	TTL time.Duration
	Immutable bool //default false
}

func New[T any]() *Cache[T] {
	db := &Cache[T]{
		storage: make(map[string]CacheStore[T]),
		mu: &sync.RWMutex{},
		TTLTicker: time.NewTicker(1 * time.Second),
		showExpiredLog: false,
		logger: os.Stdout,
	}

	go db.backgroundExpiredCleaner()

	return db
}

func (c *Cache[T]) Options() *CacheOptions {
	return &CacheOptions{}
}

// Set adds an item to the cache. If the key already exists, it overwrites the value
func (c *Cache[T]) Set(key string, value T, options ...*CacheOptions) error {

	if err := ValidateKey(key); err != nil {
		return err
	}

	if err := ValidateValue(value); err != nil {
		return err
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.Has(key) {
		if err := ValidateImmutable(c.storage[key]); err != nil {
			return err
		}
	}

	var ttl time.Time = time.Time{}
	var immutable bool = false

	if len(options) > 0 {
		for _, option := range options {
			if option.TTL > 0 {
				ttl = GetExpiration(option.TTL)
			}
			if option.Immutable {
				immutable = option.Immutable
			}
		}
	}


	c.storage[key] = CacheStore[T]{Data: value, Expiration: ttl, Immutable: immutable}
	return nil
}

// Get returns an item from the cache. If the key does not exist, it returns nil
func (c *Cache[T]) Get(key string)  (T, error) {
	var errResult T

	if err := ValidateKey(key); err != nil {
		return errResult, err
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.Has(key) {
		return errResult, fmt.Errorf(keyNotFoundErr, key)
	}

	return c.storage[key].Data, nil
}

// Delete removes an item from the cache. If the key does not exist, it returns an error
func (c *Cache[T]) Delete(key string) error {
	if err := ValidateKey(key); err != nil {
		return err
	}

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
	c.storage = make(map[string]CacheStore[T])
	c.mu.Unlock()
}

// Size returns the number of items in the cache
func (c *Cache[T]) Size() int {
	return len(c.storage)
}

// Has checks if a key exists in the cache
func (c *Cache[T]) Has(key string)  bool {
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

// Values returns a slice of all the values in the cache
func (c *Cache[T]) Values() []T {
	values := make([]T, len(c.storage))
	i := 0
	c.mu.RLock()
	for _, v := range c.storage {
		values[i] = v.Data
		i++
	}
	c.mu.RUnlock()
	return values
}
