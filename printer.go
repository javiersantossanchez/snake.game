package main

import (
	"snake/game"

	"github.com/nsf/termbox-go"
)

var (
	lastSnake []game.Cell
)

func pr(s game.Snake) {
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
