package board

type toroidalNeighborStrategy struct{}

func (s *toroidalNeighborStrategy) Neighbors(b *Board, x int, y int) []*cell {
	var neighbors []*cell

	minX := 0
	minY := 0
	maxX := b.width - 1
	maxY := b.height - 1

	for _, offset := range b.neighborCoordinateOffsets() {
		cellX := s.neighborCoord(x, offset[0], minX, maxX)
		cellY := s.neighborCoord(y, offset[1], minY, maxY)

		cell, ok := b.cell(cellX, cellY)

		if !ok { // Should not happen
			continue
		}

		neighbors = append(neighbors, cell)
	}

	return neighbors
}

func (s *toroidalNeighborStrategy) neighborCoord(coord int, offset int, minCoord int, maxCoord int) int {
	cellCoord := coord + offset

	if cellCoord < minCoord {
		cellCoord = maxCoord
	} else if cellCoord > maxCoord {
		cellCoord = minCoord
	}

	return cellCoord
}
