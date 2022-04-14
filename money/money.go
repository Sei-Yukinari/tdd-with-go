package money

type Accessor interface {
	Amount() int
	Currency() string
}

type Money struct {
	amount   int
	currency string
}

func NewMoney(amount int, currency string) Money {
	return Money{
		amount:   amount,
		currency: currency,
	}
}

func (m Money) Times(multiplier int) Expression {
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

func (m Money) Plus(a Expression) Expression {
	return NewSum(m, a)
}

func (m Money) reduce(b Bank, to string) Money {
	return NewMoney(m.amount/b.Rate(m.currency, to), to)
}

func NewDollar(amount int) Money {
	return Money{
		amount,
		"USD",
	}
}

func NewFranc(amount int) Money {
	return Money{
		amount,
		"CHF",
	}
}
