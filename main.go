package main

import (
	"fmt"
	t "github.com/nsf/termbox-go"
	"os"
	s "strconv"
	"strings"
	"time"
)

const helpMenu string = `usage:
	go-life [-Mdmn] [data inorder of flags]
example:
	go-life -Mdmn 4 1000 1 3
flags:
	-M = max amount around lifeForm, 0 =< x < 10
	-d = delay between cycles, in Milliseconds, 0 < x
	-m = min amount around lifeForm, 0 =< x < 10
	-n = min amount for new lifeForm to be created, 0 =< x < 10
`

func help() {
	t.Close()
	fmt.Println(helpMenu)
	os.Exit(3)
}

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
	maxA := 4
	minA := 1
	minC := 3
	delay := 500
	var err error = nil
	if len(os.Args) > 1 {
		flags := os.Args[1]
		if strings.ToLower(flags) == "help" || flags[0] != '-' || len(flags)-1 > len(os.Args)-2 {
			help()
		}
		for i := 1; i < len(flags); i++ {
			switch flags[i] {
			case 'M':
				maxA, err = s.Atoi(os.Args[i+1])
			case 'm':
				minA, err = s.Atoi(os.Args[i+1])
			case 'd':
				delay, err = s.Atoi(os.Args[i+1])
			case 'n':
				minC, err = s.Atoi(os.Args[i+1])
			default:
				help()
			}
		}
		if err != nil {
			help()
		}
	}
	lifeList := make(lifeFormList, 0)
	currSettings := newSettings(maxA, minA, minC, false)
	lifeList = mainCursor.placeLifeForms(lifeList)
	for {
		lifeList = cycle(currSettings, lifeList)
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}
