package main

import (
	"snake/game"
	"time"

	"github.com/nsf/termbox-go"
)

type command int

const (
	finish command = 1 + iota
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
	pr(*snake)
	var food game.Cell
	food = game.GenerateFod(snake.GetBody(), w, h)
	termbox.SetCell(food.Column, food.Row, ' ', termbox.ColorCyan, termbox.ColorCyan)

	for running := true; running; {
		select {
		case e := <-eve:
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

		case <-commandEvent:
			running = false
		default:

		}
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
				pr(*snake)
			}
			time.Sleep(80 * time.Millisecond)

		}
	}

}
