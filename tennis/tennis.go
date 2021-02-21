package tennis

type Game struct {
	score        int
	player1Score int
	player2Score int
}

func (g *Game) Score() string {
	if g.player1Score == 2 {
		return "Thirty Love"
	}
	if g.player1Score == 1 {
		return "Fifteen Love"
	}
	return "Love All"
}

func (g *Game) FirstPlayerScore() {
	g.player1Score++
}
