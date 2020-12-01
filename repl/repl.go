package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// https://tutorialedge.net/golang/improving-your-tests-with-testify-go/
type Repl struct {
	in       io.Reader
	out      io.Writer
	commands map[string]Command
}

func NewRepl() Repl {
	out := os.Stdout
	return NewReplWith(
		os.Stdin,
		out,
		map[string]Command{
			"help": NewHelp(out),
		},
	)
}

func NewReplWith(in io.Reader, out io.Writer, commands map[string]Command) Repl {
	return Repl{
		in:       in,
		out:      out,
		commands: commands,
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
