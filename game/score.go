package game

var (
	score int
)

func StartScore() {
	score = 0
}

func IncreaseScore() {
	score += 10
}

func GetCurentScore() int {
	return score
}
