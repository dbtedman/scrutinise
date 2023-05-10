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
	errorCh := make(chan error)
	var errConsole bytes.Buffer
	var outConsole bytes.Buffer
	command := cmd.VersionCommand(&resultCh)
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
	assert.Contains(t, outConsole.String(), "scrutinise version: latest")
	assert.Contains(t, outConsole.String(), "commit: n/a")
	assert.Contains(t, outConsole.String(), "built at: n/a")
}
