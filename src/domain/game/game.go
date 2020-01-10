package game

import "domain/board"

type Game struct {
	Board_ board.Board `json:"board"`
	Stats_ Stats       `json:"stats"`
}

func (g *Game) Click(x, y int) {
	tile := g.Board_.Tile(x, y)

	if tile.IsBomb() {
		g.gameLost()
	}

	if tile.IsSafe() {
		openSafeTiles := g.Board_.Tiles().SafeArea(tile).Open()
		g.Board_.Tiles().Merge(openSafeTiles)
	}

	g.Board_.UpdateTile(tile.Open())
}

func (g *Game) Board() board.Board {
	return g.Board_
}

func (g *Game) Stats() Stats {
	return g.Stats_
}

func (g *Game) gameLost() {
	g.Board_.ShowBombs()
}
