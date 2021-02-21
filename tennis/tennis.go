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
	FirstPlayerName   string
	SecondPlayerName  string
	score             int
	firstPlayerScore  int
	secondPlayerScore int
}

func (g *Game) Score() string {
	if g.IsScoresDifferent() {
		if g.IsReadyForWin() {
			if g.IsAdvantage() {
				return g.LeadingPlayer() + spaceStr + advantageStr
			}
			return g.LeadingPlayer() + spaceStr + winStr
		}
		return scoreLookup[g.firstPlayerScore] + spaceStr + scoreLookup[g.secondPlayerScore]
	}

	if g.IsDeuce() {
		return deuceStr
	}

	return scoreLookup[g.firstPlayerScore] + spaceStr + allStr
}

func (g *Game) FirstPlayerScore() {
	g.firstPlayerScore++
}

func (g *Game) SecondPlayerScore() {
	g.secondPlayerScore++
}

func (g *Game) LeadingPlayer() string {
	if g.firstPlayerScore > g.secondPlayerScore {
		return g.FirstPlayerName
	}
	return g.SecondPlayerName
}

func (g *Game) IsAdvantage() bool {
	return math.Abs(float64(g.firstPlayerScore-g.secondPlayerScore)) == 1
}

func (g *Game) IsReadyForWin() bool {
	return g.firstPlayerScore > 3 || g.secondPlayerScore > 3
}

func (g *Game) IsScoresDifferent() bool {
	return g.firstPlayerScore != g.secondPlayerScore
}

func (g *Game) IsDeuce() bool {
	return g.firstPlayerScore >= 3 && g.secondPlayerScore >= 3
}
