package main

import (
	"flag"
	"os"
	"time"

	"github.com/anthonyalberto/gameoflife/game"
)

func main() {
	pattern := flag.String("pattern",
		"glidergun",
		"initialization pattern",
	)

	width := flag.Int("width", 150, "width of the board")
	height := flag.Int("height", 50, "height of the board")
	boardType := flag.String("board", "toroidal", "board type (regular or toroidal)")
	generationTime := flag.Int("time", 50, "milli-second generation time")

	flag.Parse()

	game := game.New()
	game.Play(*width, *height, patternFile(*pattern), *boardType, time.Duration(*generationTime))
}

func patternFile(pattern string) string {
	return os.Getenv("GOPATH") + "/src/github.com/anthonyalberto/gameoflife/patterns/" + pattern + ".json"
}
