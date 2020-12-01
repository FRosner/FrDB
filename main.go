package main

import (
	"github.com/frosner/frdb/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
