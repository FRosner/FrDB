package repl

type Command interface {
	Execute(arguments string)
}
