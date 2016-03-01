package main

import (
	"flag"
	"os"

	"github.com/anthonyalberto/gameoflife/gameoflife"
)

func main() {
	patternFile := flag.String("file",
		os.Getenv("GOPATH")+"/src/github.com/anthonyalberto/gameoflife/patterns/glidergun.json",
		"path to json pattern file",
	)

	width := flag.Int("width", 150, "width of the board")
	height := flag.Int("height", 50, "height of the board")
	boardType := flag.String("board", "regular", "board type (regular or toroidal)")

	flag.Parse()

	game := gameoflife.Game{}
	game.Play(*width, *height, *patternFile, *boardType)
}
