package gameoflife

type neighborStrategy interface {
	neighbors(*board, int, int) []*cell
}
