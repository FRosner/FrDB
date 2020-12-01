package main

import (
	"github.com/frosner/frdb/repl"
	"os"
)

func main() {
	repl.NewRepl(os.Stdin, os.Stdout).Start()
}
