package main

import "github.com/anthonyalberto/gameoflife/gameoflife"

func main() {
	game := gameoflife.Game{}
	game.Play(150, 50, "patterns/glidergun.json", "regular")
}
