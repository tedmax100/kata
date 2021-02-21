package currency

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDollarMultiplication(t *testing.T) {
	five := NewDollar(5)
	assert.True(t, true, NewDollar(10).Equals(five.Times(2)))
	assert.True(t, true, NewDollar(15).Equals(five.Times(3)))
}

func TestFranceMultiplication(t *testing.T) {
	five := NewFranc(5)
	assert.True(t, true, NewFranc(10).Equals(five.Times(2)))
	assert.True(t, true, NewFranc(15).Equals(five.Times(3)))
}

func TestEquality(t *testing.T) {
	assert.True(t, true, NewDollar(5).Equals(NewDollar(5)))
	assert.False(t, false, NewDollar(5).Equals(NewDollar(6)))
}
