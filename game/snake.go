package game

var (
	previousTail Cell
)

type Snake struct {
	body []Cell
	dir  Move
}

func NewSnake() *Snake {

	bodyTmo := make([]Cell, 0)

	tailTmp := Cell{Column: 1, Row: 2}
	headTmp := Cell{Column: 2, Row: 2}
	bodyTmo = append(bodyTmo, []Cell{tailTmp, headTmp}...)
	return &Snake{body: bodyTmo, dir: Right}
}

func (s *Snake) GetHead() Cell {
	return s.body[len(s.body)-1]
}

func (s *Snake) OutRange(w int, h int) bool {
	isOutRange := false

	if s.GetHead().Column < 0 || s.GetHead().Row < 0 {
		isOutRange = true
	} else if s.GetHead().Column >= w || s.GetHead().Row >= h {
		isOutRange = true
	}
	return isOutRange
}

func (s *Snake) getTail() Cell {
	return s.body[0]
}

func (s *Snake) TurnToRight() {
	if s.dir != Left {
		s.dir = Right
	}
}

func (s *Snake) TurnToLeft() {
	if s.dir != Right {
		s.dir = Left
	}
}

func (s *Snake) TurnToUp() {
	if s.dir != Down {
		s.dir = Up
	}
}

func (s *Snake) TurnToDown() {
	if s.dir != Up {
		s.dir = Down
	}

}

func (s *Snake) GetBody() []Cell {
	return s.body
}

func (s *Snake) GrowUp() {
	s.body = append([]Cell{previousTail}, s.body...)
}

func (s *Snake) MovingForward() {
	var row int
	var column int

	var head Cell = s.GetHead()
	previousTail = s.getTail()
	switch s.dir {
	case Up:
		row = head.Row - 1
		column = head.Column
	case Down:
		row = head.Row + 1
		column = head.Column
	case Left:
		column = head.Column - 1
		row = head.Row
	case Right:
		column = head.Column + 1
		row = head.Row
	}

	s.body = s.body[1:]
	s.body = append(s.body, Cell{Column: column, Row: row})

}
