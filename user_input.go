package main

import (
	"os"
	"snake/game"

	"github.com/nsf/termbox-go"
)

func initKeyBard() (chan game.Move, chan command) {

	moveEvent := make(chan game.Move)
	commmandEvent := make(chan command)

	go listenKeyBoard(moveEvent, commmandEvent)

	return moveEvent, commmandEvent

}

func listenKeyBoard(eventMove chan game.Move, eventCommand chan command) {

	for {
		ev := termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowLeft:
				eventMove <- game.Left
			case termbox.KeyArrowDown:
				eventMove <- game.Down
			case termbox.KeyArrowRight:
				eventMove <- game.Right
			case termbox.KeyArrowUp:
				eventMove <- game.Up
			case termbox.KeyEsc:
				eventCommand <- finish
			case termbox.KeySpace:
				eventCommand <- pause
			}
		case termbox.EventError:
			os.Exit(3)
		}
	}

}
