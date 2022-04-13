package money

type Franc struct {
	Money
}

func NewFranc(amount int) *Franc {
	return &Franc{Money{amount, "Franc"}}
}

func (f *Franc) Times(multiplier int) *Franc {
	return NewFranc(f.amount * multiplier)
}
