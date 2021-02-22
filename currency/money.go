package currency

type Money struct {
	amount int
}

func (m Money) Amount() int {
	return m.amount
}
