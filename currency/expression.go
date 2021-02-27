package currency

type IExpression interface {
	Reduce(Bank, string) IMoney
	Plus(IExpression) IExpression
	Times(int) IExpression
}

type Sum struct {
	augend IExpression
	addend IExpression
}

func NewSum(augend, addend IExpression) IExpression {
	return Sum{augend: augend, addend: addend}
}

func (s Sum) Reduce(bank Bank, to string) IMoney {
	amount := s.augend.Reduce(bank, to).Amount() + s.addend.Reduce(bank, to).Amount()
	// amount := s.augend.Amount() + s.addend.Amount()
	return Money{amount: amount, currency: to}
}

func (s Sum) Plus(addend IExpression) IExpression {
	return NewSum(s, addend)
}

func (s Sum) Times(multiplier int) IExpression {
	return NewSum(s.augend.Times(multiplier), s.addend.Times(multiplier))
}
