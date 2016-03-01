package coordcollection

// CoordCollection represents a collection of cell coordinates
type CoordCollection struct {
	collection map[int]map[int]bool
}

// New builds a new CoordCollection
func New(collection [][]int) *CoordCollection {
	coordMap := map[int]map[int]bool{}
	for _, coords := range collection {
		_, exists := coordMap[coords[0]]

		if exists {
			coordMap[coords[0]][coords[1]] = true
		} else {
			coordMap[coords[0]] = map[int]bool{coords[1]: true}
		}
	}

	return &CoordCollection{coordMap}
}

// IsMember checks if the given coordinates are part of the collection.
// Implements lookup by hash table
func (c *CoordCollection) IsMember(x int, y int) bool {
	_, exists := c.collection[x][y]
	return exists
}
