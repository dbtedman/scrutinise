package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLCommand(t *testing.T) {
	// given
	resultCh := make(chan int)
	errorCh := make(chan error)
	var errConsole bytes.Buffer
	var outConsole bytes.Buffer
	command := URLCommand(&resultCh)
	command.SetErr(&errConsole)
	command.SetOut(&outConsole)

	// when
	go func() {
		errorCh <- command.Execute()
	}()
	result := <-resultCh
	err := <-errorCh

	// then
	assert.Equal(t, SuccessResult, result)
	assert.Nil(t, err)
	assert.Equal(t, "", errConsole.String())
	assert.Contains(t, outConsole.String(), "")
	assert.Contains(t, outConsole.String(), "")
}
