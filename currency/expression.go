package currency

type IExpression interface {
	Reduce(string) IMoney
}

type Sum struct {
	augend IMoney
	addend IMoney
}

func NewSum(augend, addend IMoney) IExpression {
	return Sum{augend: augend, addend: addend}
}

func (s Sum) Reduce(to string) IMoney {
	amount := s.augend.Amount() + s.addend.Amount()
	return Money{amount: amount, currency: to}
}

/* func Plus(addend IMoney) IExpression {
	return Sum
} */
