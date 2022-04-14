package money

type Expression interface {
	Plus(Expression) Expression
	reduce(bank Bank, to string) Money
}

type Sum struct {
	augend Expression
	addend Expression
}

func (s Sum) Plus(expression Expression) Expression {
	return nil
}

func NewSum(augend, addend Expression) Sum {
	return Sum{augend, addend}
}

func (s Sum) reduce(b Bank, to string) Money {
	amount := s.augend.reduce(b, to).amount + s.addend.reduce(b, to).amount
	return NewMoney(amount, to)
}
