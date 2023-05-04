package application

import "github.com/dbtedman/scrutinise/domain"

// A ScrutiniseResponseHeadersUseCase defines a series of Observations about the provided
// response headers. These Observations will, where possible, make Recommendations for
// improvements to be made.
type ScrutiniseResponseHeadersUseCase struct {
}

func (uc ScrutiniseResponseHeadersUseCase) Execute(in ScrutiniseResponseHeadersInput) (ScrutiniseResponseHeadersOutput, error) {

	anObservation, anObservationError := domain.NewObservation()

	if anObservationError != nil {
		return ScrutiniseResponseHeadersOutput{}, anObservationError
	}

	return ScrutiniseResponseHeadersOutput{
		Observations: []domain.Observation{
			anObservation,
		},
	}, nil
}

type ScrutiniseResponseHeadersInput struct {
}

type ScrutiniseResponseHeadersOutput struct {
	Observations []domain.Observation
}
