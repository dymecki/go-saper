package tile

type Tile struct {
	Position_ Position `json:"position"`
	Data_     Data     `json:"data"`
}

func (t *Tile) Empty(x, y int) Tile {
	return Value(x, y, 0)
}

func Value(x, y, value int) Tile {
	return Tile{
		Position{x, y},
		Data{Value_: value},
	}
}

func (t *Tile) Position() Position {
	return t.Position_
}

func (t *Tile) Data() Data {
	return t.Data_
}

func (t *Tile) IsBomb() bool {
	return t.Data_.Value_ == -1
}

func (t *Tile) IsNumber() bool {
	return t.Data_.Value_ > 0
}

func (t *Tile) IsSafe() bool {
	return t.Data_.Value() == 0
}

func (t *Tile) HasFlag() bool {
	return t.Data_.Flag_
}

func (t *Tile) Open() Tile {
	return Tile{t.Position_, t.Data_.WithOpen()}
}

func (t *Tile) Close() Tile {
	return Tile{t.Position_, t.Data_.WithClose()}
}

func (t *Tile) IsNull() bool {
	return t.Position_.X() == 0 && t.Position_.Y() == 0
}

func (t *Tile) ToNumber(value int) Tile {
	return Tile{t.Position_, t.Data_.WithValue(value)}
}
