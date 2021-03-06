package currency

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/* func TestDollarMultiplication(t *testing.T) {
	five := Dollar(5)
	assert.True(t, Dollar(10).Equals(five.Times(2)))
	assert.True(t, Dollar(15).Equals(five.Times(3)))
	assert.False(t, Frac(5).Equals(Frac(6)))
	assert.False(t, Frac(5).Equals(Dollar(5)))
}
*/
/* func TestFranceMultiplication(t *testing.T) {
	five := Frac(5)
	assert.True(t, Frac(10).Equals(five.Times(2)))
	assert.True(t, Frac(15).Equals(five.Times(3)))
} */

func TestEquality(t *testing.T) {
	assert.True(t, Dollar(5).Equals(Dollar(5)))
	assert.False(t, Dollar(5).Equals(Dollar(6)))
	assert.False(t, Frac(5).Equals(Dollar(5)))
}

func TestCurrency(t *testing.T) {
	assert.Equal(t, "USD", Dollar(1).Currency())
	assert.Equal(t, "CHF", Frac(1).Currency())
}

func TestMoneyString(t *testing.T) {
	assert.Equal(t, "1 USD", Dollar(1).String())
	assert.Equal(t, "1 CHF", Frac(1).String())
}

func TestSimpleAddition(t *testing.T) {
	/* sum := Dollar(5).Plus(Dollar(5))
	assert.Equal(t, Dollar(10), sum) */
	five := Dollar(5)
	sum := five.Plus(five)
	bank := NewBank()
	reduced := bank.Reduce(sum, "USD")
	assert.Equal(t, Dollar(10), reduced)
	// reduced Money := Bank.Reduce(sum, "USD")
	// assert.Equal(t, Dollar(10), reduced)
	//var five IMoney = Dollar(5)
	//sum := five.Plus(five)

}

func TestReduceSum(t *testing.T) {
	sum := NewSum(Dollar(3), Dollar(4))
	bank := NewBank()
	result := bank.Reduce(sum, "USD")
	assert.Equal(t, Dollar(7), result)
}

func TestReduceMoney(t *testing.T) {
	bank := NewBank()
	result := bank.Reduce(Dollar(1), "USD")
	assert.Equal(t, Dollar(1), result)
}

func TestReduceMoneyDifferentCurrency(t *testing.T) {
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	result := bank.Reduce(Frac(2), "USD")
	assert.Equal(t, Dollar(1), result)
}

func TestArrayEquals(t *testing.T) {
	assert.Equal(t, []string{"abc"}, []string{"abc"})
}

func TestMixedsAddition(t *testing.T) {
	var fiveBucks IExpression = Dollar(5)
	var tenFracs IExpression = Frac(10)

	var bank Bank = NewBank()
	bank.AddRate("CHF", "USD", 2)

	var result IMoney = bank.Reduce(fiveBucks.Plus(tenFracs), "USD")
	assert.Equal(t, Dollar(10), result)
}

func TestSumPlusMoney(t *testing.T) {
	var fiveBucks IExpression = Dollar(5)
	var tenFracs IExpression = Frac(10)

	var bank Bank = NewBank()
	bank.AddRate("CHF", "USD", 2)
	var sum IExpression = NewSum(fiveBucks, tenFracs).Plus(fiveBucks)
	var result IMoney = bank.Reduce(sum, "USD")
	assert.Equal(t, Dollar(15), result)
}

func TestSumTimes(t *testing.T) {
	var fiveBucks IExpression = Dollar(5)
	var tenFracs IExpression = Frac(10)

	var bank Bank = NewBank()
	bank.AddRate("CHF", "USD", 2)
	var sum IExpression = NewSum(fiveBucks, tenFracs).Times(2)
	var result IMoney = bank.Reduce(sum, "USD")
	assert.Equal(t, Dollar(20), result)
}
