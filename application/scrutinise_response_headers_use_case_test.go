package application

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanScrutiniseResponseHeaders(t *testing.T) {
	uc := ScrutiniseResponseHeadersUseCase{}
	in := ScrutiniseResponseHeadersInput{}

	result, resultError := uc.Execute(in)

	assert.Len(t, result.Observations, 1)
	assert.Nil(t, resultError)
}
