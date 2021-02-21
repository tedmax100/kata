package currency

type Franc struct {
	amount int
}

func NewFranc(amount int) Dollar {
	return Dollar{amount: amount}
}

func (f Franc) Times(multiplier int) Franc {
	return Franc{amount: f.amount * multiplier}
}

func (f Franc) Amount() int {
	return f.amount
}

func (f Franc) Equals(object interface{}) bool {
	return f.amount == object.(Franc).amount
}
