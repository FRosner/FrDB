package repl

import (
	"fmt"
	"io"
	"strings"
)

type Get struct {
	out io.Writer
}

func NewGet(out io.Writer) Get {
	return Get{
		out: out,
	}
}

func (g Get) Execute(arguments string) {
	split := strings.SplitN(arguments, " ", 2)
	if len(split) > 1 || strings.TrimSpace(split[0]) == "" {
		_, _ = fmt.Fprintln(g.out, "Illegal arguments")
	} else {
		_, _ = fmt.Fprintln(g.out, "GET", split[0])
	}

}
