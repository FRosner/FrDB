package wal

import (
	"os"
	"sync"
)

type Log struct {
	mu   sync.RWMutex
	Path string   // absolute path to log directory
	File *os.File // tail segment file handle
}

func Open(filePath string) (*Log, error) {
	log := &Log{path: filePath}

	file, err := os.OpenFile(log.File, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	} else {
		database.File = file
	}

	// TODO lock file

	// TODO initialize DB if it doesn't exist
	// TODO read existing DB, initializing pages
	// TODO memory map?

	return database, nil
}
