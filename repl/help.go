package repl

import (
	"fmt"
	"os"
)

func Help(out *os.File) {
	_, _ = fmt.Fprintln(out,
		`Available commands:
  - help: shows this text
  - exit: exits the REPL`,
	)
}
