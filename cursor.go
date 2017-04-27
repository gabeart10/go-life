package main

import (
	"github.com/nsf/termbox-go"
)

type cursor struct {
	x     int
	y     int
	color termbox.Attribute
}

var mainCursor = &cursor{
	x:     0,
	y:     0,
	color: termbox.ColorWhite,
}

func (c *cursor) placeLifeForms(list lifeFormList) lifeFormList {
	c.placeCursor(0, 0)
	for {
		event := termbox.PollEvent()
		if event.Key == termbox.KeyEnter {
			return list
		}
		switch event.Ch {
		case 'h':
			c.placeCursor(c.x-1, c.y)
		case 'j':
			c.placeCursor(c.x, c.y+1)
		case 'k':
			c.placeCursor(c.x, c.y-1)
		case 'l':
			c.placeCursor(c.x+1, c.y)
		case 'c':
			buffer := termbox.CellBuffer()
			w, _ := termbox.Size()
			if buffer[c.x+(c.y*w)].Ch != lifeFormSymbol {
				list = newLifeForm(c.x, c.y, list)
				termbox.SetCell(c.x, c.y, lifeFormSymbol, termbox.ColorBlue, c.color)
			}
			termbox.Flush()
		}
	}
	return list
}

func (c *cursor) placeCursor(x, y int) bool {
	buffer := termbox.CellBuffer()
	w, h := termbox.Size()
	if x < 0 || y < 0 || y >= h || x >= w {
		return false
	}
	termbox.SetCell(c.x, c.y, buffer[c.x+(c.y*w)].Ch, buffer[c.x+(c.y*w)].Fg, termbox.ColorDefault)
	termbox.SetCell(x, y, buffer[x+(y*w)].Ch, buffer[x+(y*w)].Fg, c.color)
	c.x = x
	c.y = y
	termbox.Flush()
	return true
}
