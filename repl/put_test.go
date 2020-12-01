package repl_test

import (
	"bytes"
	"github.com/frosner/frdb/repl"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPut_MissingArguments(t *testing.T) {
	var out bytes.Buffer
	repl.NewPut(&out).Execute("")
	assert.Equal(t, out.String(), "Illegal arguments\n")
}

func TestPut_TooFewArguments(t *testing.T) {
	var out bytes.Buffer
	repl.NewPut(&out).Execute("key")
	assert.Equal(t, out.String(), "Illegal arguments\n")
}

func TestPut(t *testing.T) {
	var out bytes.Buffer
	repl.NewPut(&out).Execute("key value")
	assert.Equal(t, out.String(), "PUT key value\n")
}
