package application

import (
	"errors"
	"net/url"
)

type ScrutiniseURLUC struct {
}

func NewScrutiniseURLUC() (*ScrutiniseURLUC, error) {
	return &ScrutiniseURLUC{}, nil
}

type ScrutiniseURLUCRequest struct {
	Url url.URL
}

type ScrutiniseURLUCResult struct {
}

func (uc ScrutiniseURLUC) Execute(request ScrutiniseURLUCRequest) (*ScrutiniseURLUCResult, error) {
	if !request.Url.IsAbs() {
		return nil, errors.New("expected absolute url")
	}

	return &ScrutiniseURLUCResult{}, nil
}
