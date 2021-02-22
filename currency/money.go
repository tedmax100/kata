package currency

import (
	"fmt"
	"reflect"
)

type IMoney interface {
	Times(multiplier int) IMoney
	Amount() int
	Currency() string
}

type Money struct {
	amount   int
	currency string
	IMoney
}

func Dollar(amount int) Money {
	return Money{amount: amount, currency: "USD"}
}

func Frac(amount int) Money {
	return Money{amount: amount, currency: "CHF"}
}

func (m Money) Times(multiplier int) Money {
	return Money{amount: m.amount * multiplier, currency: m.currency}
}

func (m Money) Equals(object interface{}) bool {
	members := reflect.ValueOf(object)
	/* absMoneyValid := members.FieldByName("Money").IsValid()
	if absMoneyValid == true {
		absMoneyField := members.FieldByName("Money")
		amount := absMoneyField.FieldByName("amount").Int()
		return m.amount == int(amount) // && //reflect.TypeOf(m) == absMoneyField.Type()
	} */
	amount := members.FieldByName("amount").Int()
	return m.amount == int(amount) && m.currency == object.(Money).currency //reflect.TypeOf(m) == reflect.TypeOf(object)
}

func (m Money) Currency() string {
	return m.currency
}

func (m Money) String() string {
	return fmt.Sprintf("%d %s", m.amount, m.currency)
}

func (m Money) Plus(addend Money) Money {
	return Money{amount: m.amount + addend.amount, currency: m.currency}
}
