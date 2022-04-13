package money

type Dollar struct {
	Money
}

func NewDollar(amount int) *Money {
	return &Money{
		amount,
		"Dollar",
		"USD",
	}
}
