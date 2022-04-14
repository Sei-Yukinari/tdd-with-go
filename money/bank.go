package money

type Bank struct {
}

func (b Bank) Reduce(source Expression, to string) Money {
	s, _ := source.(Sum)
	amount := s.augend.amount + s.addend.amount
	return Money{
		amount:   amount,
		currency: to,
	}
}
