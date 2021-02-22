package currency

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDollarMultiplication(t *testing.T) {
	five := NewDollar(5)
	assert.True(t, NewDollar(10).Equals(five.Times(2)))
	assert.True(t, NewDollar(15).Equals(five.Times(3)))
	assert.False(t, NewFranc(5).Equals(NewFranc(6)))
	// assert.False(t, NewFranc(5).Equals(NewDollar(5)))
}

func TestFranceMultiplication(t *testing.T) {
	five := NewFranc(5)
	assert.True(t, NewFranc(10).Equals(five.Times(2)))
	assert.True(t, NewFranc(15).Equals(five.Times(3)))
}

func TestEquality(t *testing.T) {
	assert.True(t, NewDollar(5).Equals(NewDollar(5)))
	assert.False(t, NewDollar(5).Equals(NewDollar(6)))
}
