package money

type Accessor interface {
	Amount() int
	Name() string
	Currency() string
}

type Money struct {
	amount   int
	name     string
	currency string
}

func (m *Money) Times(multiplier int) *Money {
	return &Money{
		amount:   m.amount * multiplier,
		name:     m.name,
		currency: m.currency,
	}
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

func (m Money) Currency() string {
	return m.currency
}

func Dollar(amount int) *Money {
	return &Money{
		amount,
		"Dollar",
		"USD",
	}
}

func Franc(amount int) *Money {
	return &Money{
		amount,
		"Franc",
		"CHF",
	}
}
