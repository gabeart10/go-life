package main

type settings struct {
	maxAround   int
	minAround   int
	minToCreate int
	countBorder bool
}

func newSettings(maxA, minA, minTC int, countB bool) *settings {
	return &settings{
		maxAround:   maxA,
		minAround:   minA,
		minToCreate: minTC,
		countBorder: countB,
	}
}
