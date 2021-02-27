package currency

import (
	"fmt"
	"reflect"
)

type IMoney interface {
	Times(int) IExpression
	Amount() int
	Currency() string
	Equals(interface{}) bool
	String() string
	Plus(IMoney) IExpression
	IExpression
	//Plus(addend IMoney) IExpression
}

type Money struct {
	amount   int
	currency string
}

func Dollar(amount int) IMoney {
	return Money{amount: amount, currency: "USD"}
}

func Frac(amount int) IMoney {
	return Money{amount: amount, currency: "CHF"}
}

func (m Money) Amount() int {
	return m.amount
}

func (m Money) Times(multiplier int) IExpression {
	return Money{amount: m.amount * multiplier, currency: m.currency}
}

func (m Money) Equals(object interface{}) bool {
	members := reflect.ValueOf(object)
	amount := members.FieldByName("amount").Int()
	return m.amount == int(amount) && m.currency == object.(Money).currency //reflect.TypeOf(m) == reflect.TypeOf(object)
}

func (m Money) Currency() string {
	return m.currency
}

func (m Money) String() string {
	return fmt.Sprintf("%d %s", m.amount, m.currency)
}

func (m Money) Plus(addend IMoney) IExpression {
	return NewSum(m, addend)
	//return Money{amount: m.amount + addend.(Money).amount, currency: m.currency}
}

func (m Money) Reduce(bank Bank, to string) IMoney {
	rate := bank.Rate(m.currency, to)

	return Money{amount: m.amount / rate, currency: to}
}

/* func (m Money) Plus(addend IMoney) IExpression {
	return NewSum(m, addend)
} */
