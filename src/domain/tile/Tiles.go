package tile

import "fmt"

type Tiles struct {
	Tiles_ map[string]Tile `json:"tiles"`
}

func (t *Tiles) Init(width, height int) Tiles {
	tiles := Tiles{}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			tiles.Put(Tile{}.Empty(x, y))
		}
	}

	return tiles
}

func (t *Tiles) FromArray(values []Tile) Tiles {
	tiles := Tiles{}

	for _, tile := range values {
		tiles.Put(tile)
	}

	return tiles
}

func (t *Tiles) Put(tile Tile) Tiles {
	t.Tiles_[tile.Position().Label()] = tile

	return *t
}

func (t *Tiles) Tile(x, y int) Tile {
	return t.Tiles_[fmt.Sprintf("%d.%d", x, y)]
}

func (t *Tiles) Position(position Position) Tile {
	return t.Tile(position.X(), position.Y())
}

func (t *Tiles) Has(tile Tile) bool {
	if _, exist := t.Tiles_[tile.Position().Label()]; exist {
		return true
	}

	return false
}

func (t *Tiles) SafeArea(tile Tile) Tiles {
	result := Tiles{}
	tmp := make([]Tiles, 1)

	for len(tmp) > 0 {
		tmp = tmp[:len(tmp)-1]

		for _, item := range tmp[len(tmp)-1].Tiles_ {
			if result.Has(item) {
				continue
			}

			if item.IsNumber() {
				result.Put(item)
			}

			if item.IsSafe() {
				result.Put(item)

				tmp = append(tmp, t.Around(item))
			}
		}
	}

	return result
}

func (t *Tiles) Open() Tiles {
	result := Tiles{}

	for _, tile := range t.Tiles_ {
		result.Put(tile.Open())
	}

	return result
}

func (t *Tiles) Close() Tiles {
	result := Tiles{}

	for _, tile := range t.Tiles_ {
		result.Put(tile.Close())
	}

	return result
}

func (t *Tiles) Numbers() Tiles {
	result := Tiles{}

	for _, tile := range t.Tiles_ {
		if tile.IsNumber() {
			result.Put(tile)
		}
	}

	return result
}

func (t *Tiles) Merge(tiles Tiles) Tiles {
	result := Tiles{}

	for _, tile := range tiles.Tiles_ {
		result.Put(tile)
	}

	return result
}

func (t *Tiles) WithNumbers() Tiles {
	result := Tiles{}

	for _, tile := range t.Tiles_ {
		result.Put(t.bombOrNumber(tile))
	}

	return result
}

func (t *Tiles) bombOrNumber(tile Tile) Tile {
	if tile.IsBomb() {
		return tile
	} else {
		return tile.ToNumber(t.Around(tile).Bombs().Count())
	}
}

func (t *Tiles) Bombs() Tiles {
	result := Tiles{}

	for _, tile := range t.Tiles_ {
		if tile.IsBomb() {
			result.Put(tile)
		}
	}

	return result
}

func (t *Tiles) Around(tile Tile) Tiles {
	x := tile.Position().X()
	y := tile.Position().Y()
	result := Tiles{}

	result.Put(t.Tile(x-1, y-1))
	result.Put(t.Tile(x, y-1))
	result.Put(t.Tile(x+1, y-1))
	result.Put(t.Tile(x-1, y))
	result.Put(t.Tile(x+1, y))
	result.Put(t.Tile(x-1, y+1))
	result.Put(t.Tile(x, y+1))
	result.Put(t.Tile(x+1, y+1))

	return result
}

func (t *Tiles) Count() int {
	return len(t.Tiles_)
}
