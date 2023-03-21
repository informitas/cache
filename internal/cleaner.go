package internal

import (
	"fmt"
	"time"
)

func (db *Cache[T]) backgroundExpiredCleaner() {
	for {
		<-db.TTLTicker.C
		db.mu.Lock()
		for key, value := range db.storage {
			if IsExpired(value.Expiration) {
				if db.showExpiredLog {
					fmt.Fprintf(db.logger, keyWasExpiredErr, key, value.Expiration.Format(time.RFC3339))
				}
				delete(db.storage, key)
			}
		}
		db.mu.Unlock()
	}
}
