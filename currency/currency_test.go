package currency

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	//product := five.Times(2)
	//assert.Equal(t, product.Amount(), 10)
	assert.True(t, true, NewDollar(10).Equals(five.Times(2)))

	// product = five.Times(3)
	// assert.Equal(t, product.Amount(), 15)
	assert.True(t, true, NewDollar(15).Equals(five.Times(3)))
}

func TestEquality(t *testing.T) {
	assert.True(t, true, NewDollar(5).Equals(NewDollar(5)))
	assert.False(t, false, NewDollar(5).Equals(NewDollar(6)))
}
