package repl_test

import (
	"github.com/frosner/frdb/repl"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestHelp(t *testing.T) {
	output, err := ioutil.TempFile("", "output-*")
	if err != nil {
		t.Fatal(err)
	}
	defer output.Close()
	defer os.Remove(output.Name())
	repl.Help(output)
	data, err := ioutil.ReadFile(output.Name())
	assert.Equal(t, string(data), `Available commands:
  - help: shows this text
  - exit: exits the REPL
`)
}
