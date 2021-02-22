package currency

type Franc struct {
	Money
}

func NewFranc(amount int) Dollar {
	return Dollar{Money{amount: amount}}
}

func (f Franc) Times(multiplier int) Franc {
	return Franc{Money{amount: f.amount * multiplier}}
}

func (f Franc) Amount() int {
	return f.amount
}
