package repl

import (
	"fmt"
	"io"
)

type Help struct {
	out io.Writer
}

func NewHelp(out io.Writer) Help {
	return Help{
		out: out,
	}
}

func (h Help) Execute(arguments string) {
	_, _ = fmt.Fprintln(h.out,
		`Available commands:
  - help:               shows this text
  - get <key>:          retrieves the value stored behind <key>
  - put <key> <value>:  stores <value> behind <key>
  - exit:               exits the REPL`,
	)
}
