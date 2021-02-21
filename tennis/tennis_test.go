package tennis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoveAll(t *testing.T) {
	game := &Game{}
	assert.Equal(t, game.Score(), "Love All")
}

func TestFifteenLove(t *testing.T) {
	game := &Game{}
	// 玩家1 先進一球
	GivenFirstPlayerScore(game, 1)
	assert.Equal(t, game.Score(), "Fifteen Love")
}

func TestThirtyLove(t *testing.T) {
	game := &Game{}
	// 玩家1 先進一球
	GivenFirstPlayerScore(game, 2)
	assert.Equal(t, game.Score(), "Thirty Love")
}

func GivenFirstPlayerScore(game *Game, times int) {
	for i := 0; i < times; i++ {
		game.FirstPlayerScore()
	}
}
