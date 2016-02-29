package gameoflife

type cell struct {
	alive     bool
	nextAlive bool
	neighbors []*cell
}

func (c *cell) computeNextState() {
	switch c.aliveNeighborsCount() {
	case 2:
		if c.alive {
			c.nextAlive = true
		} else {
			c.nextAlive = false
		}
	case 3:
		c.nextAlive = true
	default:
		c.nextAlive = false
	}
}

func (c *cell) nextGeneration() {
	c.alive = c.nextAlive
}

func (c *cell) aliveNeighborsCount() int {
	var aliveNeighborsCount int

	for _, neighbor := range c.neighbors {
		if neighbor.alive {
			aliveNeighborsCount++
		}
	}

	return aliveNeighborsCount
}
