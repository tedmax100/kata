package currency

type Bank struct{}

func (b Bank) Reduce(source IExpression, to string) IMoney {
	sum := source.(Sum)
	return sum.Reduce(to)
}
