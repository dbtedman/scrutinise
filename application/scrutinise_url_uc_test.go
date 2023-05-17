package application_test

import (
	"net/url"
	"testing"

	"github.com/dbtedman/scrutinise/application"
	"github.com/stretchr/testify/assert"
)

func TestCanScrutiniseURL(t *testing.T) {
	// given
	uc, ucError := application.NewScrutiniseURLUC()
	requestUrl, _ := url.Parse("https://example.com")
	request := application.ScrutiniseURLUCRequest{
		Url: *requestUrl,
	}

	// when
	executeResult, executeError := uc.Execute(request)

	// then
	assert.Nil(t, ucError)
	assert.Nil(t, executeError)
	assert.NotNil(t, executeResult)
}

func TestScrutiniseURLRequiresURL(t *testing.T) {
	// given
	uc, ucError := application.NewScrutiniseURLUC()
	request := application.ScrutiniseURLUCRequest{}

	// when
	executeResult, executeError := uc.Execute(request)

	// then
	assert.Nil(t, ucError)
	assert.NotNil(t, executeError)
	assert.Nil(t, executeResult)
}
