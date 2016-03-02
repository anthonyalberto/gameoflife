package board

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/anthonyalberto/gameoflife/coordcollection"
)

// Board is the game board
type Board struct {
	width  int
	height int
	neighborStrategy
	cells [][]*cell
}

// New is the factory
func New(width int, height int, aliveCoordinates *coordcollection.CoordCollection, neighborStrategyStr string) *Board {
	newBoard := Board{
		width:  width,
		height: height,
	}

	newBoard.initNeighborStrategy(neighborStrategyStr)
	newBoard.initCells(aliveCoordinates)
	newBoard.setNeighbors()

	return &newBoard
}

// Step goes to the next generation
func (b *Board) Step() {
	// TODO: use goroutines
	for i := 0; i < b.width; i++ {
		for j := 0; j < b.height; j++ {
			b.cells[i][j].computeNextState()
		}
	}

	for i := 0; i < b.width; i++ {
		for j := 0; j < b.height; j++ {
			b.cells[i][j].nextGeneration()
		}
	}
}

// Display displays the board
func (b *Board) Display() {
	// The Board output
	var outputString string

	// Scan row by row
	for i := 0; i < b.height; i++ {
		for j := 0; j < b.width; j++ {
			if b.cells[j][i].alive {
				outputString += "O"
			} else {
				outputString += " "
			}
		}

		outputString += "\n"
	}

	// Clear the screen
	clear, _ := exec.Command("clear").Output()
	os.Stdout.Write(clear)

	// Print the Board
	fmt.Printf(outputString)
}

func (b *Board) initNeighborStrategy(neighborStrategyStr string) {
	switch neighborStrategyStr {
	case "regular":
		b.neighborStrategy = &regularNeighborStrategy{}
	case "toroidal":
		b.neighborStrategy = &toroidalNeighborStrategy{}
	default:
		panic("Unsupported board type")
	}
}

func (b *Board) initCells(aliveCoordinates *coordcollection.CoordCollection) {
	for i := 0; i < b.width; i++ {
		var column []*cell

		for j := 0; j < b.height; j++ {
			alive := aliveCoordinates.IsMember(i, j)
			column = append(column, &cell{alive: alive})
		}

		b.cells = append(b.cells, column)
	}
}

func (b *Board) setNeighbors() {
	for i := 0; i < b.width; i++ {
		for j := 0; j < b.height; j++ {
			b.cells[i][j].neighbors = b.Neighbors(b, i, j)
		}
	}
}

func (b *Board) cell(x int, y int) (*cell, bool) {
	if x < 0 || y < 0 || x >= b.width || y >= b.height {
		return &cell{}, false
	}

	return b.cells[x][y], true
}

func (b *Board) neighborCoordinateOffsets() [][]int {
	return [][]int{{-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}}
}
