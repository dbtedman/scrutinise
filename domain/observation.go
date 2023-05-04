package domain

// An Observation represents the result of scrutinising something, one or more
// Recommendation will be referenced.
type Observation struct {
}

func NewObservation() (Observation, error) {
	return Observation{}, nil
}
