package db

import "os"

type database struct {
	FilePath string
	File     *os.File
}

// https://github.com/boltdb/bolt/blob/master/db.go
func Open(filePath string) (*database, error) {
	database := &database{FilePath: filePath}

	file, err := os.OpenFile(database.FilePath, os.O_RDWR|os.O_CREATE, 0600)
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

// TODO close
