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

func TestOnes(t *testing.T) {
	var rule *Rule = &Rule{}
	assert.Equal(t, 1, rule.Once([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, 0, rule.Once([]int{6, 2, 2, 4, 5}))
}
