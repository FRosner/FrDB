package repl

import (
	"fmt"
	"io"
	"strings"
)

type Put struct {
	out io.Writer
}

func NewPut(out io.Writer) Put {
	return Put{
		out: out,
	}
}

func (g Put) Execute(arguments string) {
	split := strings.SplitN(arguments, " ", 2)
	if len(split) != 2 {
		_, _ = fmt.Fprintln(g.out, "Illegal arguments")
	} else {
		_, _ = fmt.Fprintln(g.out, "PUT", split[0], split[1])
	}

}
