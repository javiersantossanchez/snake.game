package main

import (
	"snake/game"
	"time"

	"github.com/nsf/termbox-go"
)

type command int

const (
	finish command = 1 + iota
	pause
	stop
)

func crash(s game.Snake) bool {
	var head game.Cell = s.GetHead()
	for _, current := range s.GetBody()[:len(s.GetBody())-1] {
		if head.Equals(current) {
			return true
		}
	}
	return false
}

func main() {

	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	var (
		w, h = termbox.Size()
	)

	var eve, commandEvent = initKeyBard()
	snake := game.NewSnake()
	printSnake(*snake)
	var food game.Cell
	food = game.GenerateFod(snake.GetBody(), w, h)
	termbox.SetCell(food.Column, food.Row, ' ', termbox.ColorCyan, termbox.ColorCyan)

	var pauseSw = false
	for running := true; running; {
		select {
		case e := <-eve:
			if !pauseSw {
				switch e {
				case game.Left:
					snake.TurnToLeft()
				case game.Down:
					snake.TurnToDown()
				case game.Right:
					snake.TurnToRight()
				case game.Up:
					snake.TurnToUp()
				}
			}

		case event := <-commandEvent:
			switch event {
			case pause:
				pauseSw = !pauseSw
			case finish:
				running = false
			}

		default:

		}
		if !pauseSw {
			snake.MovingForward()
			if crash(*snake) {
				running = false
			} else {
				if snake.OutRange(w, h) {
					running = false
					break
				} else {

					if snake.GetHead() == food {
						snake.GrowUp()
						food = game.GenerateFod(snake.GetBody(), w, h)
						termbox.SetCell(food.Column, food.Row, ' ', termbox.ColorCyan, termbox.ColorCyan)
					}
					printSnake(*snake)
				}
				time.Sleep(80 * time.Millisecond)

			}
		}
	}

	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)
	termbox.Flush()
	printMessage(1, 1, "Game Over")

	time.Sleep(2 * time.Second)

}
