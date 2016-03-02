package board

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/anthonyalberto/gameoflife/coordcollection"
)

// Board is the game board
type Board struct {
	width  int
	height int
	neighborStrategy
	cells  [][]*cell
	output []byte
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
	newBoard.initOutput()

	return &newBoard
}

// Step goes to the next generation
func (b *Board) Step() {
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
	currentOutput := make([]byte, len(b.output))
	copy(currentOutput, b.output)

	for i := 0; i < b.width; i++ {
		for j := 0; j < b.height; j++ {
			if b.cells[i][j].alive {
				currentOutput[i+((b.width+1)*j)] = 'O'
			}
		}
	}

	// Clear the screen
	clear, _ := exec.Command("clear").Output()
	os.Stdout.Write(clear)

	// Print the Board
	fmt.Printf(string(currentOutput))
}

func (b *Board) initOutput() {
	lineString := strings.Repeat(" ", b.width) + "\n"
	wholeOutput := strings.Repeat(lineString, b.height)

	b.output = []byte(wholeOutput)
}

func (b *Board) cellCount() int {
	return b.width * b.height
}

func (b *Board) initNeighborStrategy(neighborStrategyStr string) {
	switch neighborStrategyStr {
	case "regular":
		b.neighborStrategy = &regularNeighborStrategy{}
	case "toroidal":
		b.neighborStrategy = &toroidalNeighborStrategy{}
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
