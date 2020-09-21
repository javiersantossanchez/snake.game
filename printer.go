package main

import (
	"snake/game"

	"github.com/nsf/termbox-go"
)

var (
	lastSnake []game.Cell
)

func printSnake(s game.Snake) {
	if lastSnake != nil {
		for _, body := range lastSnake {
			termbox.SetCell(body.Column, body.Row, ' ', termbox.ColorDefault, termbox.ColorDefault)

		}
	}
	for _, body := range s.GetBody() {
		termbox.SetCell(body.Column, body.Row, ' ', termbox.ColorGreen, termbox.ColorGreen)

	}
	lastSnake = s.GetBody()
	termbox.Flush()
}

func printMessage(x, y int, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, 1, c, termbox.ColorGreen, termbox.ColorDefault)
		x++
	}

	termbox.Flush()
}
