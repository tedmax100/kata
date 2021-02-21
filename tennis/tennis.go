package tennis

import "math"

var scoreLookup map[int]string

func init() {
	scoreLookup = map[int]string{
		0: "Love",
		1: "Fifteen",
		2: "Thirty",
		3: "Forty",
	}
}

type Game struct {
	Player1Name  string
	Player2Name  string
	score        int
	player1Score int
	player2Score int
}

func (g *Game) Score() string {
	// var score1, score2 string

	if g.player1Score != g.player2Score {
		if g.player1Score > 3 {
			if math.Abs(float64(g.player1Score-g.player2Score)) == 1 {
				return g.Player1Name + " Adv"
			}
		}
		return scoreLookup[g.player1Score] + " " + scoreLookup[g.player2Score]
	}

	if g.player1Score >= 3 {
		return "Deuce"
	}

	return scoreLookup[g.player1Score] + " " + "All"
}

func (g *Game) FirstPlayerScore() {
	g.player1Score++
}

func (g *Game) SecondPlayerScore() {
	g.player2Score++
}
