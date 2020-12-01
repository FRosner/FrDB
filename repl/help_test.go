package repl_test

import (
	"bytes"
	"github.com/frosner/frdb/repl"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelp(t *testing.T) {
	var out bytes.Buffer
	repl.NewHelp(&out).Execute("")
	assert.Equal(t, out.String(), `Available commands:
  - help:               shows this text
  - get <key>:          retrieves the value stored behind <key>
  - put <key> <value>:  stores <value> behind <key>
  - exit:               exits the REPL
`)
}
