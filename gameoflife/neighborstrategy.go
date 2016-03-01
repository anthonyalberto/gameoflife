package gameoflife

type neighborStrategy interface {
	Neighbors(*board, int, int) []*cell
}
