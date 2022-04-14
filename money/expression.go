package money

type Expression interface {
	Plus(Expression) Expression
	Times(int) Expression
	reduce(bank Bank, to string) Money
}

type Sum struct {
	augend Expression
	addend Expression
}

func (s Sum) Plus(e Expression) Expression {
	return NewSum(s, e)
}

func (s Sum) Times(multiplier int) Expression {
	return NewSum(s.augend.Times(multiplier), s.addend.Times(multiplier))
}

func NewSum(augend, addend Expression) Sum {
	return Sum{augend, addend}
}

func (s Sum) reduce(b Bank, to string) Money {
	amount := s.augend.reduce(b, to).amount + s.addend.reduce(b, to).amount
	return NewMoney(amount, to)
}
