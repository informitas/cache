package internal

import "io"

func (db *Cache[T]) EnableLogs() {
	db.showExpiredLog = true
}

func (db *Cache[T]) DisableLogs() {
	db.showExpiredLog = false
}

func (db *Cache[T]) SetLogger(logger io.Writer) {
	db.logger = logger
}
