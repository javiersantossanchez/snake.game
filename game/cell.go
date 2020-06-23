package game

type Cell struct {
	Row    int
	Column int
}

func (cell *Cell) Equals(other Cell) bool {
	return cell.Column == other.Column && cell.Row == other.Row
}
