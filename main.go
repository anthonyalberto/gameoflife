package main

import "github.com/anthonyalberto/gameoflife"

func main() {
	game := gameoflife.Game{}
	game.Play(50, 50, "patterns/glider.json", "regular")
}
