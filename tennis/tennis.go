package tennis

var scoreLookup map[int]string

func init() {
	scoreLookup = map[int]string{
		1: "Fifteen",
		2: "Thirty",
	}
}

type Game struct {
	score        int
	player1Score int
	player2Score int
}

func (g *Game) Score() string {
	if score, ok := scoreLookup[g.player1Score]; ok == true {
		return score + " Love"
	}
	return "Love All"
}

func (g *Game) FirstPlayerScore() {
	g.player1Score++
}
