package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiCurrencyMoney(t *testing.T) {
	t.Run("Can be multiplied by dollar", func(t *testing.T) {
		five := Dollar(5)
		assert.Equal(t, Dollar(10), five.Times(2))
		assert.Equal(t, Dollar(15), five.Times(3))
	})
	t.Run("Same amount of dollars is equivalent", func(t *testing.T) {
		assert.True(t, Dollar(5).Equals(Dollar(5)))
		assert.False(t, Dollar(5).Equals(Dollar(6)))
	})
	t.Run("Can be multiplied by franc", func(t *testing.T) {
		five := Franc(5)
		assert.Equal(t, Franc(10), five.Times(2))
		assert.Equal(t, Franc(15), five.Times(3))
	})
	t.Run("Same amount of dollar and franc are not equivalent", func(t *testing.T) {
		assert.False(t, Dollar(5).Equals(Franc(5)))
	})
	t.Run("test Currency", func(t *testing.T) {
		assert.Equal(t, "USD", Dollar(1).Currency())
		assert.Equal(t, "CHF", Franc(1).Currency())
	})
	t.Run("simple addition", func(t *testing.T) {
		five := Dollar(5)
		sum := five.Plus(five)
		reduced := Bank{}.Reduce(sum, "USD")
		assert.Equal(t, Dollar(10), reduced)
	})
	t.Run("plus returns sum", func(t *testing.T) {
		five := Dollar(10)
		result := five.Plus(five)
		s, ok := result.(Sum)
		assert.True(t, ok)
		assert.Equal(t, five, s.augend)
		assert.Equal(t, five, s.addend)
	})
	t.Run("reduce sum", func(t *testing.T) {
		sum := Bank{}.Reduce(
			NewSum(
				Dollar(3), Dollar(4),
			),
			"USD",
		)
		assert.Equal(t, Dollar(7), sum)
	})
}
