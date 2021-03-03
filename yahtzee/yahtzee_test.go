package yahtzee

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChance(t *testing.T) {
	var game *Game = NewGame()
	var n int = 20
	var pins int = 0
	rollMany(game, n, pins)
	assert.Equal(t, 0, game.Score())
}
