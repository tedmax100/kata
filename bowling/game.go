package bowling

type Game struct {
	score int
}

func NewGame() *Game {
	return &Game{
		score: 0,
	}
}

func (g *Game) Roll(pins int) {
	g.score += pins // roll職責錯誤, 它不應該去累加score
	return
}

func (g *Game) Score() int {
	return g.score //score() 反而沒負責計算分數
}
