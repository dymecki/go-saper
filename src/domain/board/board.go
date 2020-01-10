package board

import t "domain/tile"

type Board struct {
	Width_  int     `json:"width"`
	Height_ int     `json:"height"`
	Tiles_  t.Tiles `json:"tiles"`
}

func (b *Board) Empty(width, height int) Board {
	return Board{
		width,
		height,
		t.Tiles{}.Init(width, height),
	}
}

func (b *Board) CloseAll() *Board {
	b.Tiles_ = b.Tiles().Close()

	return b
}

func (b *Board) Click(x int, y int) t.Tile {
	tile := b.Tiles().Tile(x, y)
	b.Tiles_ = b.Tiles().Put(tile.Open())

	return tile
}

func (b *Board) Tile(x int, y int) t.Tile {
	return b.Tiles().Tile(x, y)
}

func (b *Board) UpdateTile(tile t.Tile) Board {
	b.Tiles_ = b.Tiles().Put(tile)

	return *b
}

func (b *Board) ShowBombs() Board {
	bombs := b.Tiles().Bombs().Open()
	b.Tiles_.Merge(bombs)

	return *b
}

func (b *Board) Width() int {
	return b.Width_
}

func (b *Board) Height() int {
	return b.Height_
}

func (b *Board) Tiles() t.Tiles {
	return b.Tiles_
}
