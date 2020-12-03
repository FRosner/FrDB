package db_test

import (
	"github.com/frosner/frdb/db"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestOpen_CreatesNewFile(t *testing.T) {
	dir, err := ioutil.TempDir("", "frdb-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	dbFilePath := dir + "/fr.db"
	db, err := db.Open(dbFilePath)
	if err != nil {
		t.Fatal(err)
	}

	assert.FileExists(t, dbFilePath)
	assert.Equal(t, db.FilePath, dbFilePath)
}

// TODO test open existing File
