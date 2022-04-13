package money

type Dollar struct {
	Money int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{amount}
}

func (d *Dollar) times(multiplier int) *Dollar {
	return NewDollar(d.Money * multiplier)
}

func (d *Dollar) equals(other *Dollar) bool {
	return d.Money == other.Money
}
