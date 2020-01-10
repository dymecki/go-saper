package game

import "domain/game"

type EasyGame struct {
	*game.AbstractGameBuilder
}

func bombs() int {
	return 9
}

func width() int {
	return 8
}

func height() int {
	return 8
}
