package board

type regularNeighborStrategy struct{}

func (s *regularNeighborStrategy) Neighbors(b *Board, x int, y int) []*cell {
	var neighbors []*cell

	for _, offset := range b.neighborCoordinateOffsets() {
		cell, ok := b.cell(x+offset[0], y+offset[1])

		if !ok {
			continue
		}

		neighbors = append(neighbors, cell)
	}

	return neighbors
}
