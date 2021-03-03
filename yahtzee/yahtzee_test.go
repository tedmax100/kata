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

func TestTwos(t *testing.T) {
	var rule *Rule = &Rule{}
	assert.Equal(t, 4, rule.Twos([]int{1, 2, 3, 2, 6}))
	assert.Equal(t, 10, rule.Twos([]int{2, 2, 2, 2, 2}))
}

func TestThrees(t *testing.T) {
	var rule *Rule = &Rule{}
	assert.Equal(t, 3, rule.Threes([]int{1, 2, 3, 2, 6}))
	assert.Equal(t, 15, rule.Threes([]int{3, 3, 3, 3, 3}))
}

func TestFours(t *testing.T) {
	var rule *Rule = &Rule{}
	assert.Equal(t, 4, rule.Fours([]int{1, 2, 3, 4, 6}))
	assert.Equal(t, 16, rule.Fours([]int{3, 4, 4, 4, 4}))
}

func TestFives(t *testing.T) {
	var rule *Rule = &Rule{}
	assert.Equal(t, 5, rule.Fives([]int{1, 2, 3, 5, 6}))
	assert.Equal(t, 20, rule.Fives([]int{3, 5, 5, 5, 5}))
}

func TestSixes(t *testing.T) {
	var rule *Rule = &Rule{}
	assert.Equal(t, 6, rule.Sixes([]int{1, 2, 3, 5, 6}))
	assert.Equal(t, 24, rule.Sixes([]int{3, 6, 6, 6, 6}))
}
