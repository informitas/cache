package internal

import "time"

type CacheStore[T any] struct {
	Data        T
	Expiration  time.Time
	Immutable		bool
}
