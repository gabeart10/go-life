package main

import (
	"github.com/nsf/termbox-go"
)

func findAmountAround(x, y int, list lifeFormList, s settings) int {
	cellBuff := termbox.CellBuffer()
	w, h := termbox.Size()
	cordsAround := make([][2]int, 8)
	amount := 0
	cordsAround = append(cordsAround, [2]int{x, y - 1})
	cordsAround = append(cordsAround, [2]int{x - 1, y - 1})
	cordsAround = append(cordsAround, [2]int{x + 1, y - 1})
	cordsAround = append(cordsAround, [2]int{x, y + 1})
	cordsAround = append(cordsAround, [2]int{x - 1, y - 1})
	cordsAround = append(cordsAround, [2]int{x + 1, y - 1})
	cordsAround = append(cordsAround, [2]int{x - 1, y})
	cordsAround = append(cordsAround, [2]int{x + 1, y})
	for _, cord := range cordsAround {
		if cord[0] < 0 || cord[1] < 0 || cord[0] > w || cord[1] > h {
			if s.countBorder == true {
				amount++
			}
		} else {
			for _, life := range list {
				if life.x == cord[0] && life.y == cord[1] {
					amount++
				}
			}
		}
	}
	return amount
}
