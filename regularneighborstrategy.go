package gameoflife

type regularNeighborStrategy struct{}

func (s *regularNeighborStrategy) neighbors(b *board, x int, y int) []*cell {
	var neighbors []*cell
	for _, offset := range s.neighborCoordinateOffsets() {
		cell, ok := b.cell(x+offset[0], y+offset[1])

		if !ok {
			continue
		}

		neighbors = append(neighbors, cell)
	}

	return neighbors
}

func (s *regularNeighborStrategy) neighborCoordinateOffsets() [][]int {
	return [][]int{{-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}}
}
