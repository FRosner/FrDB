package repl

import (
	"fmt"
	"io"
)

func Help(out io.Writer) {
	_, _ = fmt.Fprintln(out,
		`Available commands:
  - help: shows this text
  - exit: exits the REPL`,
	)
}
