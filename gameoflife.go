package main

import (
	"flag"
	"os"
	"time"

	"github.com/anthonyalberto/gameoflife/game"
)

func main() {
	patternFile := flag.String("file",
		os.Getenv("GOPATH")+"/src/github.com/anthonyalberto/gameoflife/patterns/glidergun.json",
		"path to json pattern file",
	)

	width := flag.Int("width", 150, "width of the board")
	height := flag.Int("height", 50, "height of the board")
	boardType := flag.String("board", "regular", "board type (regular or toroidal)")
	generationTime := flag.Int("time", 50, "milli-second generation time")

	flag.Parse()

	game := game.New()
	game.Play(*width, *height, *patternFile, *boardType, time.Duration(*generationTime))
}
