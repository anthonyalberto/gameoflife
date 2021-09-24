package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/anthonyalberto/gameoflife/game"
)

//go:embed patterns
var embededFiles embed.FS

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
	game.Play(*width, *height, patternFileContent(*pattern), *boardType, time.Duration(*generationTime))
}

func patternFileContent(pattern string) []byte {
	data, err := embededFiles.ReadFile(fmt.Sprintf("patterns/%s.json", pattern))

	if err != nil {
		log.Fatal(err)
	}

	return data
}
