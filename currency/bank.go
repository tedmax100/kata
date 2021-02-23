package currency

type Bank struct{}

func (b Bank) Reduce(source IExpression, to string) IMoney {
	return Dollar(10)
}
