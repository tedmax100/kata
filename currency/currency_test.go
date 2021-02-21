package currency

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	five.Times(2)

	assert.Equal(t, five.Amount(), 10)
}

type Dollar struct {
	amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{amount: amount}
}

func (d *Dollar) Times(multiplier int) {
	d.amount *= multiplier
}

func (d *Dollar) Amount() int {
	return d.amount
}
