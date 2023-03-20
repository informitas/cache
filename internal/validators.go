package internal

import (
	"errors"
	"fmt"
	"time"
)

func ValidateKey(key string) error {
	if len(key) > 0 {
		return nil
	}

	return errors.New(keyEmptyErr)
}

func ValidateValue(value any) error {
	if value != nil {
		return nil
	}

	return errors.New(valueEmptyErr)
}

func GetKeyNotFoundError(key string) error {
	return fmt.Errorf(keyNotFoundErr, key)
}

func FormatDuration(d time.Duration, format string) string {
	return time.Unix(0, 0).UTC().Add(d).Format(format)
}

func ValidateImmutable[T any](store CacheStore[T]) error {
	if store.Immutable {
		return errors.New(immutableErr)
	}

	return nil
}
