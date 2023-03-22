package cache

import (
	"github.com/informitas/cache/internal"
)

// New creates a new cache
func New[T any]() *internal.Cache[T] {
	return internal.New[T]()
}

// Options returns the cache options
func Options() *internal.CacheOptions {
	return &internal.CacheOptions{}
}
