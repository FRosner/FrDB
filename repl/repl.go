package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

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

func Start(in io.Reader, out io.Writer) {
	reader := bufio.NewReader(in)
	Help(out)
	printPrompt(out)
	command := readCommand(reader)
	for ; !isExit(command); command = readCommand(reader) {
		switch command {
		case "help":
			Help(out)
		default:
			printInvalidCommand(out, command)
		}
		printPrompt(out)
	}
	_, _ = fmt.Fprintln(out, "Bye!")
}
