package currency

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	product := five.Times(2)
	assert.Equal(t, product.Amount(), 10)

	product = five.Times(3)
	assert.Equal(t, product.Amount(), 15)
}

type Dollar struct {
	amount int
}

func NewDollar(amount int) Dollar {
	return Dollar{amount: amount}
}

func (d Dollar) Times(multiplier int) Dollar {
	return Dollar{amount: d.amount * multiplier}
}

func (d Dollar) Amount() int {
	return d.amount
}
