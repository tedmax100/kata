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
	// 玩家1 先進1球
	GivenFirstPlayerScore(game, 1)
	assert.Equal(t, game.Score(), "Fifteen Love")
}

func TestThirtyLove(t *testing.T) {
	game := &Game{}
	// 玩家1 先進2球
	GivenFirstPlayerScore(game, 2)
	assert.Equal(t, game.Score(), "Thirty Love")
}

func TestFortyLove(t *testing.T) {
	game := &Game{}
	// 玩家1 先進3球
	GivenFirstPlayerScore(game, 3)
	assert.Equal(t, game.Score(), "Forty Love")
}

func TestLoveFifteen(t *testing.T) {
	game := &Game{}
	// 玩家2 先進1球
	GivenSecondPlayerScore(game, 1)
	assert.Equal(t, game.Score(), "Love Fifteen")
}

func TestLoveThirty(t *testing.T) {
	game := &Game{}
	// 玩家2 先進2球
	GivenSecondPlayerScore(game, 2)
	assert.Equal(t, game.Score(), "Love Thirty")
}

func TestLoveForty(t *testing.T) {
	game := &Game{}
	// 玩家2 先進2球

	GivenSecondPlayerScore(game, 3)
	assert.Equal(t, game.Score(), "Love Forty")
}

func TestFifteenAll(t *testing.T) {
	game := &Game{}
	// 玩家1&2 各進1
	GivenFirstPlayerScore(game, 1)
	GivenSecondPlayerScore(game, 1)
	assert.Equal(t, game.Score(), "Fifteen All")
}
func TestThirtyAll(t *testing.T) {
	game := &Game{}
	// 玩家1&2 各進1
	GivenFirstPlayerScore(game, 2)
	GivenSecondPlayerScore(game, 2)
	assert.Equal(t, game.Score(), "Thirty All")
}
func TestDeuce(t *testing.T) {
	game := &Game{}
	// 玩家1&2 各進3
	GivenFirstPlayerScore(game, 3)
	GivenSecondPlayerScore(game, 3)
	assert.Equal(t, game.Score(), "Deuce")
}

func TestDeuceWhen_4_4(t *testing.T) {
	game := &Game{}
	// 玩家1&2 各進4
	GivenFirstPlayerScore(game, 4)
	GivenSecondPlayerScore(game, 4)
	assert.Equal(t, game.Score(), "Deuce")
}

func TestFirstPlayerAdvantage(t *testing.T) {
	game := &Game{Player1Name: "Nathan", Player2Name: "Able"}
	// 玩家1進4球, 玩家2進了3球
	GivenFirstPlayerScore(game, 4)
	GivenSecondPlayerScore(game, 3)
	assert.Equal(t, game.Score(), "Nathan Adv")
}

func TestSecondPlayerAdvantage(t *testing.T) {
	game := &Game{Player1Name: "Nathan", Player2Name: "Able"}
	// 玩家1進3球, 玩家2進了4球
	GivenFirstPlayerScore(game, 3)
	GivenSecondPlayerScore(game, 4)
	assert.Equal(t, game.Score(), "Able Adv")
}

func GivenFirstPlayerScore(game *Game, times int) {
	for i := 0; i < times; i++ {
		game.FirstPlayerScore()
	}
}

func GivenSecondPlayerScore(game *Game, times int) {
	for i := 0; i < times; i++ {
		game.SecondPlayerScore()
	}
}
