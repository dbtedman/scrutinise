package application

import "github.com/dbtedman/scrutinise/domain"

type ScrutiniseResponseHeadersUseCase struct {
}

type ScrutiniseResponseHeadersInput struct {
}

type ScrutiniseResponseHeadersOutput struct {
	Observations []domain.Observation
}

func (uc ScrutiniseResponseHeadersUseCase) Execute(in ScrutiniseResponseHeadersInput) (ScrutiniseResponseHeadersOutput, error) {
	return ScrutiniseResponseHeadersOutput{}, nil
}
