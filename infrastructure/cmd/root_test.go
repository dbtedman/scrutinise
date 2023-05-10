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
	errorCh := make(chan error)
	var errConsole bytes.Buffer
	var outConsole bytes.Buffer
	command := cmd.RootCommand(&resultCh)
	command.SetErr(&errConsole)
	command.SetOut(&outConsole)

	// when
	go func() {
		errorCh <- command.Execute()
	}()
	result := <-resultCh
	err := <-errorCh

	// then
	assert.Equal(t, cmd.SuccessResult, result)
	assert.Nil(t, err)
	assert.Equal(t, "", errConsole.String())
	assert.Contains(t, outConsole.String(), "Tool to scrutinise website development security")
	assert.Contains(t, outConsole.String(), "-h, --help   help for scrutinise")
}
