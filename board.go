package gameoflife

type board struct {
	width            int
	height           int
	neighborStrategy neighborStrategy
	cells            [][]cell
}

func (b *board) New(width int, height int, aliveCoordinates coordCollection, neighborStrategyStr string) {
	b.width = width
	b.height = height
	b.initNeighborStrategy(neighborStrategyStr)

	b.initCells(aliveCoordinates)
	b.setNeighbors()
}

func (b *board) initNeighborStrategy(neighborStrategyStr string) {
	switch neighborStrategyStr {
	case "regular":
		b.neighborStrategy = &regularNeighborStrategy{}
		//case "tororial":
	}
}

func (b *board) initCells(aliveCoordinates coordCollection) {
	for i := 0; i < b.width; i++ {
		var column []cell

		for j := 0; j < b.height; j++ {
			alive := aliveCoordinates.isMember(i, j)
			column = append(column, cell{alive: alive})
		}

		b.cells = append(b.cells, column)
	}
}

func (b *board) setNeighbors() {
	for i := 0; i < b.width; i++ {
		for j := 0; j < b.height; j++ {
			b.cells[i][j].neighbors = b.neighborStrategy.neighbors(b, i, j)
		}
	}
}
