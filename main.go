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
	lifeList := make(lifeFormList, 5)
	currSettings := newSettings(4, 2, 3, false)
	for i := 0; i < 5; i++ {
		newLifeForm(30, 20+i, lifeList)
	}
	for {
		lifeList = cycle(currSettings, lifeList)
		time.Sleep(300 * time.Millisecond)
	}
}
