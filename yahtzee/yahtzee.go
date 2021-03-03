package yahtzee

type Game struct {
}

// Create a Game instance
func NewGame() *Game {
	return &Game{}
}

// Throw() : given a game having 5 dice with random number between 1 and 6

// ReThrow(idxList []int) : given index between 1 and 5, will re-throw them;
// 然後就發現...測試不了XDD, 因為是隨機數字;
// 但想了想,我幹麻把規則跟Game綁在一起呢?

func (g *Game) Chance()
