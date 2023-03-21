package internal

import "time"

const (
	keyEmptyErr    = "key is empty"
	keyNotFoundErr = "key not found, key = %q"

	valueEmptyErr = "value is empty"
	immutableErr  = "value is immutable"

	keyWasExpiredErr = "key was expired [key = %q, expiration = %q]\n"
)

func GetNow() time.Time {
	return time.Now()
}

func GetExpiration(ttl time.Duration) time.Time {
	return GetNow().Add(ttl)
}

func IsExpired(timer time.Time) bool {
	if timer.IsZero() {
		return false
	}
	return GetNow().After(timer)
}
