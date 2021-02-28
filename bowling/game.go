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
		if g.isStrike(frameIdx) {
			score += 10 + g.getStrikeBonus(frameIdx)
			frameIdx++
		} else if g.isSpare(frameIdx) {
			score += 10 + g.getSpareBonus(frameIdx)
			frameIdx += 2
		} else {
			score += g.sumOfBallsInFrame(frameIdx)
			frameIdx += 2
		}
	}
	return score
}

func (g *Game) isStrike(frameIdx int) bool {
	return g.rolls[frameIdx] == 10
}

func (g *Game) sumOfBallsInFrame(frameIdx int) int {
	return g.rolls[frameIdx] + g.rolls[frameIdx+1]
}

func (g *Game) getSpareBonus(frameIdx int) int {
	return g.rolls[frameIdx+2]
}

func (g *Game) getStrikeBonus(frameIdx int) int {
	return g.rolls[frameIdx+1] + g.rolls[frameIdx+2]
}

func (g *Game) isSpare(frameIdx int) bool {
	return g.rolls[frameIdx]+g.rolls[frameIdx+1] == 10
}
