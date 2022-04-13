package money

type Dollar struct {
	Money
}

func NewDollar(amount int) *Dollar {
	return &Dollar{Money{amount, "Dollar"}}
}

func (d *Dollar) Times(multiplier int) *Dollar {
	return NewDollar(d.amount * multiplier)
}
