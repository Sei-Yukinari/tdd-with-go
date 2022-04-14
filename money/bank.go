package money

type Bank struct {
	rates map[CurrencyMap]int
}

type CurrencyMap [2]string

func NewBank() Bank {
	return Bank{rates: make(map[CurrencyMap]int)}
}

func (b Bank) Reduce(source Expression, to string) Money {
	return source.reduce(b, to)
}

func (b Bank) AddRate(from, to string, rate int) {
	b.rates[CurrencyMap{from, to}] = rate
	return
}

func (b Bank) Rate(from, to string) (rate int) {
	if from == to {
		return 1
	}

	if r, ok := b.rates[CurrencyMap{from, to}]; ok {
		rate = r
	}
	return rate
}
