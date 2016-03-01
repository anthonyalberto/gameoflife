package board

type neighborStrategy interface {
	Neighbors(*Board, int, int) []*cell
}
