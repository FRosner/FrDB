package main

import (
	"github.com/frosner/frdb/repl"
	"io"
	"os"
)

func main() {
	var in io.Reader = os.Stdin
	repl.Start(in, os.Stdout)
}
