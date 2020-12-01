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
			"get":  NewGet(out),
			"put":  NewPut(out),
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

func readCommand(r *bufio.Reader) (command string, arguments string) {
	raw, _ := r.ReadString('\n')
	trimmed := strings.TrimSpace(raw)
	split := strings.SplitN(trimmed, " ", 2)
	if len(split) == 1 {
		return split[0], ""
	} else {
		return split[0], split[1]
	}
}

func isExit(command string) bool {
	return strings.EqualFold("exit", command)
}

func (r Repl) Start() {
	reader := bufio.NewReader(r.in)
	printPrompt(r.out)
	commandText, commandArguments := readCommand(reader)
	for ; !isExit(commandText); commandText, commandArguments = readCommand(reader) {
		if command, commandExists := r.commands[commandText]; commandExists {
			command.Execute(commandArguments)
		} else {
			printInvalidCommand(r.out, commandText)
		}
		printPrompt(r.out)
	}
	_, _ = fmt.Fprintln(r.out, "Bye!")
}
