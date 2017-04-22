package main

import (
	t "github.com/nsf/termbox-go"
	"os"
	"time"
)

func main() {
	t.Init()
	go func() {
		for {
			var event = t.PollEvent()
			if event.Key == t.KeyCtrlC {
				t.Close()
				os.Exit(3)
			}
		}
	}()
	lifeList := make(lifeFormList, 0)
	currSettings := newSettings(4, 1, 3, false)
	lifeList = mainCursor.placeLifeForms(lifeList)
	for {
		lifeList = cycle(currSettings, lifeList)
		time.Sleep(500 * time.Millisecond)
	}
}
