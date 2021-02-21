package tennis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoveAll(t *testing.T) {
	game := NewGame("Nathan", "Able")
	assert.Equal(t, game.Score(), "Love All")
}

func TestFifteenLove(t *testing.T) {
	game := NewGame("Nathan", "Able")
	GivenFirstPlayerScore(game, 1)
	assert.Equal(t, game.Score(), "Fifteen Love")
}

func TestThirtyLove(t *testing.T) {
	game := NewGame("Nathan", "Able")
	GivenFirstPlayerScore(game, 2)
	assert.Equal(t, game.Score(), "Thirty Love")
}

func TestFortyLove(t *testing.T) {
	game := NewGame("Nathan", "Able")
	GivenFirstPlayerScore(game, 3)
	assert.Equal(t, game.Score(), "Forty Love")
}

func TestLoveFifteen(t *testing.T) {
	game := NewGame("Nathan", "Able")
	GivenSecondPlayerScore(game, 1)
	assert.Equal(t, game.Score(), "Love Fifteen")
}

func TestLoveThirty(t *testing.T) {
	game := NewGame("Nathan", "Able")
	GivenSecondPlayerScore(game, 2)
	assert.Equal(t, game.Score(), "Love Thirty")
}

func TestLoveForty(t *testing.T) {
	game := NewGame("Nathan", "Able")
	GivenSecondPlayerScore(game, 3)
	assert.Equal(t, game.Score(), "Love Forty")
}

func TestFifteenAll(t *testing.T) {
	game := NewGame("Nathan", "Able")
	GivenFirstPlayerScore(game, 1)
	GivenSecondPlayerScore(game, 1)
	assert.Equal(t, game.Score(), "Fifteen All")
}
func TestThirtyAll(t *testing.T) {
	game := NewGame("Nathan", "Able")
	GivenFirstPlayerScore(game, 2)
	GivenSecondPlayerScore(game, 2)
	assert.Equal(t, game.Score(), "Thirty All")
}
func TestDeuce(t *testing.T) {
	game := NewGame("Nathan", "Able")
	GivenFirstPlayerScore(game, 3)
	GivenSecondPlayerScore(game, 3)
	assert.Equal(t, game.Score(), "Deuce")
}

func TestDeuceWhen_4_4(t *testing.T) {
	game := NewGame("Nathan", "Able")
	GivenFirstPlayerScore(game, 4)
	GivenSecondPlayerScore(game, 4)
	assert.Equal(t, game.Score(), "Deuce")
}

func TestFirstPlayerAdvantage(t *testing.T) {
	game := NewGame("Nathan", "Able")
	GivenFirstPlayerScore(game, 4)
	GivenSecondPlayerScore(game, 3)
	assert.Equal(t, game.Score(), "Nathan Adv")
}

func TestSecondPlayerAdvantage(t *testing.T) {
	game := NewGame("Nathan", "Able")
	GivenFirstPlayerScore(game, 3)
	GivenSecondPlayerScore(game, 4)
	assert.Equal(t, game.Score(), "Able Adv")
}

func TestFirstPlayerWinTheGame(t *testing.T) {
	game := NewGame("Nathan", "Able")
	GivenFirstPlayerScore(game, 4)
	GivenSecondPlayerScore(game, 2)
	assert.Equal(t, game.Score(), "Nathan Win")
}

func TestSecondPlayerWinTheGame(t *testing.T) {
	game := NewGame("Nathan", "Able")
	GivenFirstPlayerScore(game, 2)
	GivenSecondPlayerScore(game, 4)
	assert.Equal(t, game.Score(), "Able Win")
}

func GivenFirstPlayerScore(game *Game, times int) {
	for i := 0; i < times; i++ {
		game.FirstPlayerWinTheBall()
	}
}

func GivenSecondPlayerScore(game *Game, times int) {
	for i := 0; i < times; i++ {
		game.SecondPlayerWinTheBall()
	}
}
