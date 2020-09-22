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
	start
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

	boardWidth := w - 2
	boardHeight := h - 3

	var eve, commandEvent = initKeyBard()

	var finishsw = false
	for !finishsw {
		printWalls(boardWidth, boardHeight)
		game.StartScore()
		printScore(game.GetCurentScore())
		printInstructions()

		snake := game.NewSnake()
		printSnake(*snake)
		var food game.Cell
		food = game.GenerateFod(snake.GetBody(), boardWidth, boardHeight)
		printFood(food)

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
					if snake.OutRange(boardWidth, boardHeight) {
						running = false
						break
					} else {

						if snake.GetHead() == food {
							snake.GrowUp()
							game.IncreaseScore()
							printScore(game.GetCurentScore())
							food = game.GenerateFod(snake.GetBody(), boardWidth, boardHeight)
							printFood(food)
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
		printMessage(10, 10, "Do you want play again? Y/N")

		finishsw = true
		select {
		case events := <-commandEvent:
			switch events {
			case start:
				finishsw = false
			default:
				finishsw = true
			}
		case e := <-eve:
			switch e {
			default:
				finishsw = true
			}
		}

		//time.Sleep(800 * time.Millisecond)
		termbox.Clear(coldef, coldef)
	}

}
