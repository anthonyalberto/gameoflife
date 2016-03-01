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
	fileContent := p.readFile()

	var aliveCoords aliveCoordinates
	err := json.Unmarshal(fileContent, &aliveCoords)
	checkFileError(err)

	return &coordCollection{aliveCoords.Alive}
}

func (p *patternParser) readFile() []byte {
	data, err := ioutil.ReadFile(p.patternFilePath)
	checkFileError(err)

	return data
}

func checkFileError(e error) {
	if e != nil {
		panic(e)
	}
}
