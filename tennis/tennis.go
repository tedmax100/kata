package tennis

var scoreLookup map[int]string

func init() {
	scoreLookup = map[int]string{
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
	if g.player2Score > 0 {
		return "Love " + scoreLookup[g.player2Score]
	}
	/* if g.player2Score == 2 {
		return "Love Thirty"
	} */
	if g.player2Score == 1 {
		return "Love Fifteen"
	}
	if g.player1Score > 0 {
		score1, _ := scoreLookup[g.player1Score]
		return score1 + " Love"

	}

	return "Love All"
}

func (g *Game) FirstPlayerScore() {
	g.player1Score++
}

func (g *Game) SecondPlayerScore() {
	g.player2Score++
}
