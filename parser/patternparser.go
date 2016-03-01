package parser

import (
	"encoding/json"
	"io/ioutil"

	"github.com/anthonyalberto/gameoflife/coordcollection"
)

// PatternParser extracts alive starting coordinates from a json file
type PatternParser struct {
	patternFilePath string
}

type aliveCoordinates struct {
	Alive [][]int `json:"alive"`
}

// New builds a new parser
func New(patternFilePath string) *PatternParser {
	return &PatternParser{patternFilePath}
}

// ExtractCoordinates extracts coordinates from the file
func (p *PatternParser) ExtractCoordinates() *coordcollection.CoordCollection {
	fileContent := p.readFile()

	var aliveCoords aliveCoordinates
	err := json.Unmarshal(fileContent, &aliveCoords)
	checkFileError(err)

	return coordcollection.New(aliveCoords.Alive)
}

func (p *PatternParser) readFile() []byte {
	data, err := ioutil.ReadFile(p.patternFilePath)
	checkFileError(err)

	return data
}

func checkFileError(e error) {
	if e != nil {
		panic(e)
	}
}
