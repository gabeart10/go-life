package main

import (
	"github.com/nsf/termbox-go"
)

func findAmountAround(x, y int, list lifeFormList, s *settings) (int, bool) {
	w, h := termbox.Size()
	cordsAround := make([][2]int, 8)
	amount := 0
	cordsAround[0] = [2]int{x, y - 1}
	cordsAround[1] = [2]int{x - 1, y - 1}
	cordsAround[2] = [2]int{x + 1, y - 1}
	cordsAround[3] = [2]int{x, y + 1}
	cordsAround[4] = [2]int{x - 1, y + 1}
	cordsAround[5] = [2]int{x + 1, y + 1}
	cordsAround[6] = [2]int{x - 1, y}
	cordsAround[7] = [2]int{x + 1, y}
	isALifeForm := false
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
				if life.x == x && life.y == y {
					isALifeForm = true
				}
			}
		}
	}
	return amount, isALifeForm
}

func cycle(s *settings, list lifeFormList) lifeFormList {
	w, h := termbox.Size()
	var toCreate = make([][2]int, 0)
	dataChan := make(chan [2]int, w*h)
	doneChan := make(chan bool, h)
	amountDone := 0
	for y := 0; y < h; y++ {
		go func(data chan [2]int, done chan bool) {
			for x := 0; x < w; x++ {
				amount, isALifeForm := findAmountAround(x, y, list, s)
				if isALifeForm == true {
					if amount >= s.minAround && amount < s.maxAround {
						dataChan <- [2]int{x, y}
					}
				} else {
					if amount >= s.minToCreate {
						dataChan <- [2]int{x, y}
					}
				}
			}
			done <- true
		}(dataChan, doneChan)
	}
	for amountDone != h {
		select {
		case data := <-dataChan:
			toCreate = append(toCreate, data)
		case <-doneChan:
			amountDone++
		}
	}
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	tempList := make(lifeFormList, len(toCreate))
	for _, cord := range toCreate {
		newLifeForm(cord[0], cord[1], tempList)
		termbox.SetCell(cord[0], cord[1], lifeFormSymbol, termbox.ColorBlue, termbox.ColorDefault)
	}
	termbox.Flush()
	return tempList
}
