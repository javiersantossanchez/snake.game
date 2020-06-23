package game

import (
	"math/rand"
	"time"
)

func GenerateFod(snakeBody []Cell, w int, h int) Cell {
	var food Cell
	for foodOnValidPosition := false; !foodOnValidPosition; {
		rand.Seed(time.Now().UnixNano())
		var column int = rand.Intn(w)
		var row int = rand.Intn(h)
		food = Cell{Row: row, Column: column}
		foodOnValidPosition = true
		for _, body := range snakeBody {
			if body.Row == food.Row && body.Column == food.Column {
				foodOnValidPosition = false
				break
			}

		}
	}
	return food
}
