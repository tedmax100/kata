package yahtzee

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChance(t *testing.T) {
	var rule *Rule = &Rule{}
	assert.Equal(t, 15, rule.Chance([]int{2, 3, 4, 5, 1}))
	assert.Equal(t, 16, rule.Chance([]int{3, 3, 4, 5, 1}))
}
