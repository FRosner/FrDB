package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func printPrompt(out *os.File) {
	_, _ = fmt.Fprint(out, "FrDB> ")
}

func printInvalidCommand(out *os.File, command string) {
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

func Start(in *os.File, out *os.File) {
	commands := map[string]interface{}{
		"help": Help,
	}
	reader := bufio.NewReader(in)
	Help(out)
	printPrompt(out)
	text := readCommand(reader)
	for ; !isExit(text); text = readCommand(reader) {
		if command, commandExists := commands[text]; commandExists {
			command.(func())()
		} else {
			printInvalidCommand(out, text)
		}
		printPrompt(out)
	}
	_, _ = fmt.Fprintln(out, "Bye!")
}
