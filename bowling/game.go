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
	var frameIdx int = 0 // i is a bad name for this variable
	for frame := 0; frame < 10; frame++ {
		if g.rolls[frameIdx]+g.rolls[frameIdx+1] == 10 { //spare; ugly comment in conditional
			score += 10 + g.rolls[frameIdx+2]
			frameIdx += 2
		} else {

			score += g.rolls[frameIdx] + g.rolls[frameIdx+1]
			frameIdx += 2
		}

	}

	return score
}
