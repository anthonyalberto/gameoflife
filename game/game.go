package game

import (
	"time"

	"github.com/anthonyalberto/gameoflife/board"
	"github.com/anthonyalberto/gameoflife/parser"
)

// Game is the point of entry of the whole game
type Game struct {
	*board.Board
}

// New creates a new Game
func New() *Game {
	return &Game{}
}

// Play is the entry point to start a new game
func (g *Game) Play(boardWidth int, boardHeight int, patternFileContent []byte, neighborStrategy string, generationTime time.Duration) {
	parser := parser.New(patternFileContent)

	g.Board = board.New(boardWidth, boardHeight, parser.ExtractCoordinates(), neighborStrategy)

	g.start(generationTime)
}

func (g *Game) start(generationTime time.Duration) {
	for {
		g.Display()
		time.Sleep(generationTime * time.Millisecond)
		g.Step()
	}
}
