package repl_test

import (
	"bytes"
	"github.com/frosner/frdb/repl"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGet_MissingArguments(t *testing.T) {
	var out bytes.Buffer
	repl.NewGet(&out).Execute("")
	assert.Equal(t, out.String(), "Illegal arguments\n")
}

func TestGet_TooManyArguments(t *testing.T) {
	var out bytes.Buffer
	repl.NewGet(&out).Execute("key arg2")
	assert.Equal(t, out.String(), "Illegal arguments\n")
}

func TestGet(t *testing.T) {
	var out bytes.Buffer
	repl.NewGet(&out).Execute("key")
	assert.Equal(t, out.String(), "GET key\n")
}
