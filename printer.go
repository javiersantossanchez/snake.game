package main

import (
	"snake/game"
	"strconv"

	"github.com/nsf/termbox-go"
)

var (
	lastSnake []game.Cell
)

func printScore(score int) {
	x := 1
	var title string = "Score:"
	for _, c := range title {
		termbox.SetCell(x, 0, c, termbox.ColorGreen, termbox.ColorDefault)
		x++
	}

	x += 2
	for _, c := range strconv.Itoa(score) {
		termbox.SetCell(x, 0, c, termbox.ColorGreen, termbox.ColorDefault)
		x++
	}

	termbox.Flush()
}

func printInstructions() {
	x := 20
	var text string = "Use arrows to move snake         TAB : Pause | ESC : Exit"
	for _, c := range text {
		termbox.SetCell(x, 0, c, termbox.ColorGreen, termbox.ColorDefault)
		x++
	}
}

func printSnake(s game.Snake) {
	if lastSnake != nil {
		for _, body := range lastSnake {
			termbox.SetCell(body.Column+1, body.Row+2, ' ', termbox.ColorDefault, termbox.ColorDefault)

		}
	}
	for _, body := range s.GetBody() {
		termbox.SetCell(body.Column+1, body.Row+2, ' ', termbox.ColorGreen, termbox.ColorGreen)

	}
	lastSnake = s.GetBody()
	termbox.Flush()
}

func printFood(food game.Cell) {
	termbox.SetCell(food.Column+1, food.Row+2, ' ', termbox.ColorRed, termbox.ColorRed)
}

func printMessage(x, y int, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, 1, c, termbox.ColorGreen, termbox.ColorDefault)
		x++
	}

	termbox.Flush()
}

/* ColorBlack
ColorRed
ColorGreen
ColorYellow
ColorBlue
ColorMagenta
ColorCyan
ColorWhite */

func printWalls(w int, h int) {
	for i := 1; i < h+3; i++ {
		termbox.SetCell(0, i, ' ', termbox.ColorCyan, termbox.ColorCyan) // |
		//termbox.SetCell(1, i, ' ', termbox.ColorGreen, termbox.ColorGreen)
	}

	for i := 1; i < h+3; i++ {
		termbox.SetCell(w+1, i, ' ', termbox.ColorCyan, termbox.ColorCyan) //|
		//termbox.SetCell(9, i, ' ', termbox.ColorGreen, termbox.ColorGreen)
	}

	for i := 0; i < w+2; i++ {
		termbox.SetCell(i, 1, ' ', termbox.ColorCyan, termbox.ColorCyan) // --
	}

	for i := 0; i < w+2; i++ {
		termbox.SetCell(i, h+2, ' ', termbox.ColorCyan, termbox.ColorCyan) // --
	}

}
