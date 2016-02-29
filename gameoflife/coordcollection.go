package gameoflife

type coordCollection struct {
	collection [][]int
}

func (c *coordCollection) isMember(x int, y int) bool {
	for _, coords := range c.collection {
		if coords[0] == x && coords[1] == y {
			return true
		}
	}

	return false
}
