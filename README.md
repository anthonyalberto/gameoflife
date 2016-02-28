# Game of Life
Go implementation of Conway's Game of Life

**Under development**

## Goals
- Outputs to terminal
- Implement cell state computation using goroutines
- Provide an interface for the Board. Implement different boards (regular, toroidal)
- Provide a parser that reads initial state from file
- Provide common pattern files
- Options will include : height and width of the board, board type, initial state file


## Architecture

### Cell
- Build a Cell type that has a `state`, `NextState`, and `Neighbors`. `Neighbors` will be a slice that contains pointers to the cell's neighbors
- Implements a `nextState()` that will determine, according to neighbors, the next generation state
- Implements `GoToNext()` that will set `state` as `nextState`. `nextState` becomes nil

### CellCoordinates
- Just a wrapper for x,y coordinates of a given Cell

### baseBoard
- Provides the methods that are shared across the different board types.
- Has a `height`, a `width` and a collection of `cells`
- `cells` is a two dimensional array that represents the board tiles
- `New(width, height, initialState)` will implement the logic of creating all the cells and store them in the board. Will also need to set the Neighbors on each cell
- provides `Display()` that will render the current state of the board in the console

### LifeBoard
- `LifeBoard` will be an interface. Types that implement it will need to provide `neighbors(CellCoordinates)`.

### RegularLifeBoard
- Has a `baseBoard`
- Will implement `neighbors` without wrapping around the left most tiles with the right most tiles

### ToroidalLifeBoard
- Has a `baseBoard`
- Will implement `neighbors` wrapping around the left most tiles with the right most tiles

### Game
- `Game` will be a type that provides `Play()`.
- Next generation computation will be triggered by `next()`
- `next()` will :
  - call a goroutine for each cell that will call `cell.NextState()`
  - wait for those goroutines to be done
  - call a goroutine for each cell that will call `cell.GoToNext()`
  - wait for those goroutines to be done
  - Display() the board

### InitialStateParser
- Will take a json file that contains the initial state of the board and convert it to an array of `CellCoordinates` that are alive
- Structure of json files :
  ```json
    "alive": [
      [0, 0],
      [0, 1],
      [0, 2]
    ]
  ```
