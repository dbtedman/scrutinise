package cmd_test

import (
	"bytes"
	"testing"

	"github.com/dbtedman/scrutinise/infrastructure/cmd"
	"github.com/stretchr/testify/assert"
)

func TestVersionCommand(t *testing.T) {
	// given
	resultCh := make(chan int)
	var errConsole bytes.Buffer
	var outConsole bytes.Buffer
	var versionCommandExecuteError error
	command := cmd.VersionCommand(&resultCh)
	command.SetErr(&errConsole)
	command.SetOut(&outConsole)

	// when
	go func() {
		versionCommandExecuteError = command.Execute()
	}()
	versionCommandResult := <-resultCh

	// then
	assert.Equal(t, cmd.SuccessResult, versionCommandResult)
	assert.Nil(t, versionCommandExecuteError)
	assert.Equal(t, "", errConsole.String())
	assert.Contains(t, outConsole.String(), "scrutinise version: latest")
	assert.Contains(t, outConsole.String(), "commit: n/a")
	assert.Contains(t, outConsole.String(), "built at: n/a")
}
