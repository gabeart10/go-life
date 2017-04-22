package main

type settings struct {
	maxAround   int
	minAround   int
	toCreate    int
	countBorder bool
}

func newSettings(maxA, minA, minTC int, countB bool) *settings {
	return &settings{
		maxAround:   maxA,
		minAround:   minA,
		toCreate:    minTC,
		countBorder: countB,
	}
}
