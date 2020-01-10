package tile

import (
	"fmt"
)

type Position struct {
	X_ int `json:"x"`
	Y_ int `json:"y"`
}

func (p *Position) X() int {
	return p.X_
}

func (p *Position) Y() int {
	return p.Y_
}

func (p *Position) Same(p2 *Position) bool {
	return p.X_ == p2.X_ && p.Y_ == p2.Y_
}

func (p *Position) Label() string {
	return fmt.Sprintf("%d.%d", p.X_, p.Y_)
}
