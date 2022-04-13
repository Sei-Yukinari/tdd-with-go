package money

type Accessor interface {
	Amount() int
	Name() string
}

type Money struct {
	amount int
	name   string
}

func (m *Money) Times(multiplier int) *Money {
	return &Money{amount: m.amount * multiplier, name: m.name}
}

func (m Money) Equals(a Accessor) bool {
	return m.Amount() == a.Amount() && m.Name() == a.Name()
}

func (m Money) Amount() int {
	return m.amount
}

func (m Money) Name() string {
	return m.name
}
