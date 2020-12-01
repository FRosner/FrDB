package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Repl struct {
	in       io.Reader
	out      io.Writer
	commands map[string]Command
}

func NewRepl(in io.Reader, out io.Writer) Repl {
	return Repl{
		in:  in,
		out: out,
		commands: map[string]Command{
			"help": NewHelp(out),
		},
	}
}

func printPrompt(out io.Writer) {
	_, _ = fmt.Fprint(out, "FrDB> ")
}

func printInvalidCommand(out io.Writer, command string) {
	trimmedCommand := strings.TrimSuffix(command, "\n")
	if trimmedCommand != "" {
		_, _ = fmt.Fprintln(out, "Unknown command: "+trimmedCommand)
	}
}

func readCommand(r *bufio.Reader) string {
	t, _ := r.ReadString('\n')
	return strings.TrimSpace(t)
}

func isExit(command string) bool {
	return strings.EqualFold("exit", command)
}

func (r Repl) Start() {
	reader := bufio.NewReader(r.in)
	r.commands["help"].Execute("")
	printPrompt(r.out)
	commandText := readCommand(reader)
	for ; !isExit(commandText); commandText = readCommand(reader) {
		if command, commandExists := r.commands[commandText]; commandExists {
			command.Execute("") // FIXME split in command and arguments
		} else {
			printInvalidCommand(r.out, commandText)
		}
		printPrompt(r.out)
	}
	_, _ = fmt.Fprintln(r.out, "Bye!")
}
