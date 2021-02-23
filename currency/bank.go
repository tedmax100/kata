package currency

type Bank struct {
	rates map[Pair]int
}

func NewBank() Bank {
	return Bank{rates: make(map[Pair]int)}
}

func (b Bank) Reduce(source IExpression, to string) IMoney {
	/* 	if reflect.TypeOf(source).Name() == "Money" {
		return source.(Money)
	} */
	//sum := source.(Sum)
	return source.Reduce(b, to) //sum.Reduce(to)
}

func (b Bank) Rate(from, to string) int {
	if from == to {
		return 1
	}
	return b.rates[NewPair(from, to)]
}

func (b Bank) AddRate(from, to string, rate int) {
	b.rates[NewPair(from, to)] = rate
}
