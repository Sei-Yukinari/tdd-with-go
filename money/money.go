package money

type Accessor interface {
	Amount() int
	Currency() string
}

type Money struct {
	amount   int
	currency string
}

func (m Money) Times(multiplier int) Money {
	return Money{
		amount:   m.amount * multiplier,
		currency: m.currency,
	}
}

func (m Money) Equals(a Accessor) bool {
	return m.Amount() == a.Amount() && m.Currency() == a.Currency()
}

func (m Money) Amount() int {
	return m.amount
}

func (m Money) Currency() string {
	return m.currency
}

func (m Money) Plus(a Money) Expression {
	return NewSum(m, a)
}

func Dollar(amount int) Money {
	return Money{
		amount,
		"USD",
	}
}

func Franc(amount int) Money {
	return Money{
		amount,
		"CHF",
	}
}
