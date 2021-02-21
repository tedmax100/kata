package tennis

import "math"

// declare const variables
var (
	scoreLookup  map[int]string
	spaceStr     string = " "
	advantageStr string = "Adv"
	winStr       string = "Win"
	deuceStr     string = "Deuce"
	allStr       string = "All"
)

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
	if g.IsScoresDifferent() {
		if g.IsReadyForWin() {
			if g.IsAdvantage() {
				return g.AdvPlayer() + spaceStr + advantageStr
			}
			return g.AdvPlayer() + spaceStr + winStr
		}
		return scoreLookup[g.player1Score] + spaceStr + scoreLookup[g.player2Score]
	}

	if g.player1Score >= 3 {
		return deuceStr
	}

	return scoreLookup[g.player1Score] + spaceStr + allStr
}

func (g *Game) FirstPlayerScore() {
	g.player1Score++
}

func (g *Game) SecondPlayerScore() {
	g.player2Score++
}

func (g *Game) AdvPlayer() string {
	if g.player1Score > g.player2Score {
		return g.Player1Name
	}
	return g.Player2Name
}

func (g *Game) IsAdvantage() bool {
	return math.Abs(float64(g.player1Score-g.player2Score)) == 1
}

func (g *Game) IsReadyForWin() bool {
	return g.player1Score > 3 || g.player2Score > 3
}

func (g *Game) IsScoresDifferent() bool {
	return g.player1Score != g.player2Score
}
