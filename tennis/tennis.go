package tennis

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
	score        int
	player1Score int
	player2Score int
}

func (g *Game) Score() string {
	// var score1, score2 string
	if g.player1Score != g.player2Score {
		return scoreLookup[g.player1Score] + " " + scoreLookup[g.player2Score]
	}
	return "Love All"
}

func (g *Game) FirstPlayerScore() {
	g.player1Score++
}

func (g *Game) SecondPlayerScore() {
	g.player2Score++
}
