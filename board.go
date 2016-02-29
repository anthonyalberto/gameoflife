package gameoflife

import (
	"fmt"
	"os"
	"os/exec"
)

type board struct {
	width            int
	height           int
	neighborStrategy neighborStrategy
	cells            [][]*cell
}

func newBoard(width int, height int, aliveCoordinates *coordCollection, neighborStrategyStr string) *board {
	newBoard := board{
		width:  width,
		height: height,
	}
	newBoard.initNeighborStrategy(neighborStrategyStr)

	newBoard.initCells(aliveCoordinates)
	newBoard.setNeighbors()

	return &newBoard
}

func (b *board) initNeighborStrategy(neighborStrategyStr string) {
	switch neighborStrategyStr {
	case "regular":
		b.neighborStrategy = &regularNeighborStrategy{}
		//case "tororial":
	}
}

func (b *board) initCells(aliveCoordinates *coordCollection) {
	for i := 0; i < b.width; i++ {
		var column []*cell

		for j := 0; j < b.height; j++ {
			alive := aliveCoordinates.isMember(i, j)
			column = append(column, &cell{alive: alive})
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

func (b *board) cell(x int, y int) (*cell, bool) {
	if x < 0 || y < 0 || x >= b.width || y >= b.height {
		return &cell{}, false
	}

	return b.cells[x][y], true
}

func (b *board) step() {
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

func (b *board) display() {
	var outputString string

	for i := 0; i < b.width; i++ {
		for j := 0; j < b.height; j++ {
			if b.cells[i][j].alive {
				outputString += "O"
			} else {
				outputString += " "
			}
		}

		outputString += "\n"
	}

	clear, _ := exec.Command("clear").Output()
	os.Stdout.Write(clear)
	// os.Stdout.Write(outputString)
	fmt.Printf(outputString)
}
