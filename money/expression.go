package money

type Expression interface {
}

type Sum struct {
	augend Money
	addend Money
}

func NewSum(augend, addend Money) Sum {
	return Sum{augend, addend}
}
