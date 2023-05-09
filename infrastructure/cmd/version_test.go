package cmd_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/dbtedman/scrutinise/infrastructure/cmd"
	"github.com/stretchr/testify/assert"
)

func TestVersionCommand(t *testing.T) {
	resultCh := make(chan int)
	var errConsole bytes.Buffer
	var outConsole bytes.Buffer
	go func() {
		command := cmd.VersionCommand(&resultCh)
		command.SetErr(&errConsole)
		command.SetOut(&outConsole)

		err := command.Execute()

		if err != nil {
			fmt.Println(err)
			resultCh <- cmd.ErrorResult
		}
	}()
	if result := <-resultCh; true {
		assert.Equal(t, "", errConsole.String())
		assert.Equal(t, "scrutinise version: latest, commit: n/a, built at: n/a\n", outConsole.String())
		assert.Equal(t, cmd.SuccessResult, result)
	} else {
		t.Error("Unexpected behaviour")
	}
}
