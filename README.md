# Game of Life
Go implementation of Conway's Game of Life

## Installation

```sh
  go get github.com/anthonyalberto/gameoflife
```

## Usage

Provided `$GOPATH/bin` is in your `$PATH` :

```sh
  gameoflife [options]
```

Supported options :
- `-board`: `regular` or `toroidal`. Toroidal wraps the board around. Default: `regular`
- `-time`: how many ms to pause between generations. Default: 30
- `-pattern`: pattern name. Look into the `patterns` folder for all supported values. Default: `glidergun`
- `-width`: the width of the board. Adapt to your terminal. Default: `150`
- `-height`: the height of the board. Adapt to your terminal. Default: `50`

### Example

Play a game that:
- has the `glidergun` pattern
- has a toroidal board that is 140x54
- pauses between generations for 40ms

```sh
  gameoflife -board=toroidal -time=40 -pattern=glidergun -width=140 -height=54
```
