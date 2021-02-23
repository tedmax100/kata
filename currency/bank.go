package currency

type Bank struct{}

func (b Bank) Reduce(source IExpression, to string) IMoney {
	/* 	if reflect.TypeOf(source).Name() == "Money" {
		return source.(Money)
	} */
	//sum := source.(Sum)
	return source.Reduce(to) //sum.Reduce(to)
}
