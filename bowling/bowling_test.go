package bowling

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGutterGame(t *testing.T) {
	var game Game = NewGame()
	for i := 0; i < 20; i++ {
		game.Roll(0)
	}
	assert.Equal(t, 0, game.Score())
}
