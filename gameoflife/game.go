package gameoflife

import "time"

// Game is the point of entry of the whole game
type Game struct {
	*board
}

// Play is the entry point to start a new game
func (g *Game) Play(boardWidth int, boardHeight int, patternFilePath string, neighborStrategy string, generationTime time.Duration) {
	parser := patternParser{patternFilePath}

	g.board = newBoard(boardWidth, boardHeight, parser.extractCoordinates(), neighborStrategy)

	g.start(generationTime)
}

func (g *Game) start(generationTime time.Duration) {
	for {
		g.Display()
		time.Sleep(generationTime * time.Millisecond)
		g.Step()
	}
}
