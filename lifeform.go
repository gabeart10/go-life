package main

import (
	"errors"
)

const (
	lifeFormSymbol rune = '🍆'
)

type lifeForm struct {
	x int
	y int
}

func newLifeForm(x, y int, list lifeFormList) *lifeForm {
	tempLife := &lifeForm{
		x: x,
		y: y,
	}
	for i := 0; i < len(list); i++ {
		if list[i] == nil {
			list[i] = tempLife
			return tempLife
		}
	}
	return nil
}

func (l *lifeForm) removeLifeForm(list lifeFormList) error {
	for i := 0; i < len(list); i++ {
		if list[i] == l {
			list = append(list[:i], list[i+1:]...)
			return nil
		}
	}
	return errors.New("removeLifeForm: not found in list")
}
