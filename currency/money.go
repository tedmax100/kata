package currency

import (
	"reflect"
)

type Money interface {
	Times(multiplier int) Money
	Amount() int
	Currency() string
}

type AbstractMoney struct {
	amount   int
	currency string
	Money
}

/* func (m Money) Amount() int {
	return m.amount
} */

/* func NewDollar(amount int) Dollar {
	return Dollar{Money{amount: amount}}
}
*/
func (m AbstractMoney) Times(multiplier int) AbstractMoney {
	return AbstractMoney{amount: m.amount * multiplier}
}

func (m AbstractMoney) Equals(object interface{}) bool {
	members := reflect.ValueOf(object)
	// moneyField := members.FieldByName("Money")
	/* moneyField := members.FieldByName("Money")
	amount := moneyField.FieldByName("amount").Int() */
	amount := members.FieldByName("amount").Int()
	return m.amount == int(amount) && reflect.TypeOf(m) == reflect.TypeOf(object)
}

func (m AbstractMoney) Currency() string {
	return m.currency
}
