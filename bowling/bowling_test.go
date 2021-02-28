package bowling

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGutterGame(t *testing.T) {
	var game *Game = NewGame()
	var n int = 20
	var pins int = 0
	rollMany(game, n, pins)
	assert.Equal(t, 0, game.Score())
}

func TestAllOnes(t *testing.T) {
	var game *Game = NewGame()
	var n int = 20
	var pins int = 1
	rollMany(game, n, pins)
	assert.Equal(t, 20, game.Score())
}

func rollMany(game *Game, n, pins int) {
	for i := 0; i < n; i++ {
		game.Roll(pins)
	}
}
