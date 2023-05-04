package application

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanScrutiniseResponseHeaders(t *testing.T) {
	uc := ScrutiniseResponseHeadersUseCase{}
	in := ScrutiniseResponseHeadersInput{}

	result, resultError := uc.Execute(in)

	assert.NotNil(t, result)
	assert.Nil(t, resultError)
}
