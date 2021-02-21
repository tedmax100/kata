package tennis

type Player struct {
	name  string
	score int
}

func NewPlayer(name string) *Player {
	return &Player{name: name}
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetScore() int {
	return p.score
}

func (p *Player) WinTheBall() {
	p.score++
}
