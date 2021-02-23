package currency

type Bank struct{}

func (b Bank) Reduce(source IExpression, to string) IMoney {
	/* 	if reflect.TypeOf(source).Name() == "Money" {
		return source.(Money)
	} */
	//sum := source.(Sum)
	return source.Reduce(b, to) //sum.Reduce(to)
}

func (b Bank) Rate(from, to string) int {
	if from == "CHF" {
		return 2
	}
	return 1
}
