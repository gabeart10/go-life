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

func (c *cursor) placeLifeForms() bool {
	return true
}

func (c *cursor) placeCursor(x, y int) bool {
	buffer := termbox.CellBuffer()
	w, h := termbox.Size()
	if x < 0 || y < 0 || y > h || x > w {
		return false
	}
	termbox.SetCell(c.x, c.y, buffer[c.x+(c.y*w)].Ch, buffer[c.x+(c.y*w)].Fg, termbox.ColorDefault)
	termbox.SetCell(x, y, buffer[x+(y*w)].Ch, buffer[x+(y*w)].Fg, c.color)
	c.x = x
	c.y = y
	return true
}
