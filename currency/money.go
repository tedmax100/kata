package currency

import (
	"reflect"
)

type Money struct {
	amount int
}

func (m Money) Amount() int {
	return m.amount
}

func (m Money) Equals(object interface{}) bool {
	members := reflect.ValueOf(object)
	moneyField := members.FieldByName("Money")
	amount := moneyField.FieldByName("amount").Int()
	return m.amount == int(amount)
}
