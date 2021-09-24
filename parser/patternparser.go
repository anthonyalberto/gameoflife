package parser

import (
	"encoding/json"

	"github.com/anthonyalberto/gameoflife/coordcollection"
)

// PatternParser extracts alive starting coordinates from a json file
type PatternParser struct {
	patternFileContent []byte
}

type aliveCoordinates struct {
	Alive [][]int `json:"alive"`
}

// New builds a new parser
func New(patternFileContent []byte) *PatternParser {
	return &PatternParser{patternFileContent}
}

// ExtractCoordinates extracts coordinates from the file
func (p *PatternParser) ExtractCoordinates() *coordcollection.CoordCollection {
	var aliveCoords aliveCoordinates
	err := json.Unmarshal(p.patternFileContent, &aliveCoords)
	checkFileError(err)

	return coordcollection.New(aliveCoords.Alive)
}

func checkFileError(e error) {
	if e != nil {
		panic(e)
	}
}
