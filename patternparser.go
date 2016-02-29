package gameoflife

import (
	"encoding/json"
	"io/ioutil"
)

type patternParser struct {
	patternFilePath string
}

type aliveCoordinates struct {
	Alive [][]int `json:"alive"`
}

func (p *patternParser) extractCoordinates() *coordCollection {
	a := coordCollection{}
	fileContent := p.readFile()

	var aliveCoords aliveCoordinates
	err := json.Unmarshal(fileContent, &aliveCoords)
	check(err)

	return &a
}

func (p *patternParser) readFile() []byte {
	data, err := ioutil.ReadFile(p.patternFilePath)
	check(err)

	return data
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
