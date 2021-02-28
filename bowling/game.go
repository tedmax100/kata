package bowling

type Game struct {
	score       int
	rolls       []int // to store maximum 22 pins of rolls
	currentRoll int
}

func NewGame() *Game {
	return &Game{
		score:       0,
		rolls:       make([]int, 21),
		currentRoll: 0,
	}
}

func (g *Game) Roll(pins int) {
	g.rolls[g.currentRoll] = pins
	g.currentRoll++
}

func (g *Game) Score() int {
	var score int = 0
	var i int = 0
	for frame := 0; frame < 10; frame++ {
		if g.rolls[i]+g.rolls[i+1] == 10 { //spare
			score += 10 + g.rolls[i+2]
			i += 2
		} else {

			score += g.rolls[i] + g.rolls[i+1]
			i += 2
		}

	}

	return score
}
