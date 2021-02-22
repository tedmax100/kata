package currency

type Dollar struct {
	AbstractMoney
}

func NewDollar(amount int) AbstractMoney {
	return AbstractMoney{amount: amount, currency: "USD"}
}

/* func (d Dollar) Times(multiplier int) Dollar {
	return Dollar{Money{amount: d.amount * multiplier}}
}
*/
/* func (d Dollar) Amount() int {
	return d.amount
}
*/
