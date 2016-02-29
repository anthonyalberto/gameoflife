package gameoflife

import (
	"fmt"
	"time"
)

// Game is the point of entry of the whole game
type Game struct {
	gameBoard *board
}

// Play is the entry point to start a new game
func (g *Game) Play(boardWidth int, boardHeight int, patternFilePath string, neighborStrategy string) {
	parser := patternParser{patternFilePath: patternFilePath}

	g.gameBoard = newBoard(boardWidth, boardHeight, parser.extractCoordinates(), neighborStrategy)

	g.start()
}

func (g *Game) start() {
	for {
		fmt.Printf("YO")
		g.gameBoard.display()
		time.Sleep(500 * time.Millisecond)
		g.gameBoard.step()
	}
}
