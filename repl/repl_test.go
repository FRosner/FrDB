package repl_test

import (
	"bytes"
	"github.com/frosner/frdb/repl"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockCommand struct {
	mock.Mock
}

func (m *MockCommand) Execute(arguments string) {
	m.Called(arguments)
}

func TestReplExecutesCommandAndExits_CommandWithoutArgs(t *testing.T) {
	var in bytes.Buffer
	var out bytes.Buffer
	mockCommand := new(MockCommand)
	commands := map[string]repl.Command{
		"cmd1": mockCommand,
	}
	repl := repl.NewReplWith(&in, &out, commands)

	mockCommand.On("Execute", "").Return()
	in.WriteString("cmd1\n")
	in.WriteString("exit\n")
	repl.Start()
	mockCommand.AssertCalled(t, "Execute", "")
}

func TestReplExecutesCommandAndExits_CommandWithArgs(t *testing.T) {
	var in bytes.Buffer
	var out bytes.Buffer
	mockCommand := new(MockCommand)
	commands := map[string]repl.Command{
		"cmd1": mockCommand,
	}
	repl := repl.NewReplWith(&in, &out, commands)

	mockCommand.On("Execute", "arg1 arg2").Return()
	in.WriteString("cmd1 arg1 arg2\n")
	in.WriteString("exit\n")
	repl.Start()
	mockCommand.AssertCalled(t, "Execute", "arg1 arg2")
}

func TestReplHandlesUnknownCommands(t *testing.T) {
	var in bytes.Buffer
	var out bytes.Buffer
	commands := map[string]repl.Command{}
	repl := repl.NewReplWith(&in, &out, commands)

	in.WriteString("unknown\n")
	in.WriteString("exit\n")
	repl.Start()
	assert.Matches(t, out.String(), "FrDB> Unknown command: unknown")
}
