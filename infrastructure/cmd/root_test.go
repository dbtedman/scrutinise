package cmd_test

import (
	"bytes"
	"testing"

	"github.com/dbtedman/scrutinise/infrastructure/cmd"
	"github.com/stretchr/testify/assert"
)

func TestRootCommand(t *testing.T) {
	// given
	resultCh := make(chan int)
	var errConsole bytes.Buffer
	var outConsole bytes.Buffer
	var rootCommandExecuteError error
	command := cmd.RootCommand(&resultCh)
	command.SetErr(&errConsole)
	command.SetOut(&outConsole)

	// when
	go func() {
		rootCommandExecuteError = command.Execute()
	}()
	rootCommandResult := <-resultCh

	// then
	assert.Equal(t, cmd.SuccessResult, rootCommandResult)
	assert.Nil(t, rootCommandExecuteError)
	assert.Equal(t, "", errConsole.String())
	assert.Contains(t, outConsole.String(), "Tool to scrutinise website development security")
	assert.Contains(t, outConsole.String(), "-h, --help   help for scrutinise")
}
