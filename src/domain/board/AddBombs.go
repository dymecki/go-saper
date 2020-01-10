package board

import (
	"domain/tile"
	"math/rand"
	"time"
)

type AddBombs struct {
	board Board
	bombs int
}

func (a *AddBombs) Board() Board {
	return Board{
		a.board.Width(),
		a.board.Height(),
		a.tiles().WithNumbers(),
	}
}

func (a *AddBombs) tiles() tile.Tiles {
	tiles := tile.Tiles{}

	for y, row := range a.matrix() {
		for x, value := range row {
			tiles.Put(tile.Value(x, y, value))
		}
	}

	return tiles
}

func (a *AddBombs) matrix() [][]int {
	values := a.values()
	matrix := [][]int{}

	for i := 0; i < a.board.Width(); i++ {
		matrix[i] = values[i*a.board.Width() : (i+1)*a.board.Width()]
	}

	return matrix
}

func (a *AddBombs) values() []int {
	result := make([]int, a.board.Tiles().Count())

	for i := 0; i < a.bombs; i++ {
		result[i] = -1
	}

	return a.randomOrder(
		append(result, make([]int, a.nonBombTilesAmount())...),
	)
}

// func (a *AddBombs) values_() []int {
// 	bombs := make([]int, a.bombs)
//
// 	for i := range bombs {
// 		bombs[i] = -1
// 	}
//
// 	empty := make([]int, a.nonBombTilesAmount())
//
// 	result := make([]int, a.board.Tiles().Count())
// 	result = append(result, bombs...)
// 	result = append(result, empty...)
//
// 	return result
// }

func (a *AddBombs) nonBombTilesAmount() int {
	return a.board.Tiles().Count() - a.bombs
}

func (a *AddBombs) randomOrder(values []int) []int {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(values), func(i, j int) {
		values[i], values[j] = values[j], values[i]
	})

	return values
}
