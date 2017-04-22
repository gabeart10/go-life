package main

import (
	"errors"
)

const (
	lifeFormSymbol rune = 'X'
)

type lifeForm struct {
	x int
	y int
}

func newLifeForm(x, y int, list lifeFormList) lifeFormList {
	tempLife := &lifeForm{
		x: x,
		y: y,
	}
	return append(list, tempLife)
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
