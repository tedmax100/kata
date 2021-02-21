package tennis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoveAll(t *testing.T) {
	game := &Game{}
	assert.Equal(t, game.Score(), "LoveAll")
}

func TestFifteenLove(t *testing.T) {
	game := &Game{}
	// 玩家1 先進一球
	game.FirstPlayerScore()
	assert.Equal(t, game.Score(), "FifteenLove")
}

func TestThirtyLove(t *testing.T) {
	game := &Game{}
	// 玩家1 先進一球
	game.FirstPlayerScore()
	game.FirstPlayerScore()
	assert.Equal(t, game.Score(), "FifteenLove")
}
