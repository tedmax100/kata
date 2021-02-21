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
	firstPlayer  *Player
	seconfPlayer *Player
}

func NewGame(firstPlayerName, secondPlayerName string) *Game {
	return &Game{
		firstPlayer:  NewPlayer(firstPlayerName),
		seconfPlayer: NewPlayer(secondPlayerName),
	}
}

func (g *Game) Score() string {
	if g.IsScoresDifferent() {
		if g.IsReadyForWin() {
			if g.IsAdvantage() {
				return g.LeadingPlayer() + spaceStr + advantageStr
			}
			return g.LeadingPlayer() + spaceStr + winStr
		}
		return scoreLookup[g.firstPlayer.GetScore()] + spaceStr + scoreLookup[g.seconfPlayer.GetScore()]
	}

	if g.IsDeuce() {
		return deuceStr
	}

	return scoreLookup[g.firstPlayer.GetScore()] + spaceStr + allStr
}

func (g *Game) FirstPlayerWinTheBall() {
	g.firstPlayer.WinTheBall()
}

func (g *Game) SecondPlayerWinTheBall() {
	g.seconfPlayer.WinTheBall()
}

func (g *Game) LeadingPlayer() string {
	if g.firstPlayer.GetScore() > g.seconfPlayer.GetScore() {
		return g.firstPlayer.GetName()
	}
	return g.seconfPlayer.GetName()
}

func (g *Game) IsAdvantage() bool {
	return math.Abs(float64(g.firstPlayer.GetScore()-g.seconfPlayer.GetScore())) == 1
}

func (g *Game) IsReadyForWin() bool {
	return g.firstPlayer.GetScore() > 3 || g.seconfPlayer.GetScore() > 3
}

func (g *Game) IsScoresDifferent() bool {
	return g.firstPlayer.GetScore() != g.seconfPlayer.GetScore()
}

func (g *Game) IsDeuce() bool {
	return g.firstPlayer.GetScore() >= 3 && g.seconfPlayer.GetScore() >= 3 && g.firstPlayer.GetScore() == g.seconfPlayer.GetScore()
}
