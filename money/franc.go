package money

type Franc struct {
	Money
}

func NewFranc(amount int) *Money {
	return &Money{amount, "Franc"}
}
