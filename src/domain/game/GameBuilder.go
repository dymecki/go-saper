package game

type GameBuilder interface {
	make() Game
	bombs() int
	width() int
	height() int
}

type AbstractGameBuilder struct {
	GameBuilder
}

func make() {
	// board := board.Board{}.Empty()

}
