package bowling

type Game struct {
}

func NewGame() Game {
	return Game{}
}

func (g Game) Roll(pins int) {
	return
}

func (g Game) Score() int {
	return 0
}
