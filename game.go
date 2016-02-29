package gameoflife

// Game is the point of entry of the whole game
type Game struct {
	board *board
}

// Play is the entry point to start a new game
func (g *Game) Play(boardWidth int, boardHeight int, patternFilePath string, neighborStrategy string) {
	parser := patternParser{patternFilePath: patternFilePath}

	g.board = newBoard(boardWidth, boardHeight, parser.extractCoordinates(), neighborStrategy)

	g.start()
}

func (g *Game) start() {
	for {

	}
}
