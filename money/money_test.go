package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiCurrencyMoney(t *testing.T) {
	t.Run("Can be multiplied by dollar", func(t *testing.T) {
		five := NewDollar(5)
		assert.Equal(t, NewDollar(10), five.Times(2))
		assert.Equal(t, NewDollar(15), five.Times(3))
	})
	t.Run("Same amount of dollars is equivalent", func(t *testing.T) {
		assert.True(t, NewDollar(5).Equals(NewDollar(5)))
		assert.False(t, NewDollar(5).Equals(NewDollar(6)))
	})
	t.Run("Can be multiplied by franc", func(t *testing.T) {
		five := NewFranc(5)
		assert.Equal(t, NewFranc(10), five.Times(2))
		assert.Equal(t, NewFranc(15), five.Times(3))
	})
	t.Run("Same amount of dollar and franc are not equivalent", func(t *testing.T) {
		assert.False(t, NewDollar(5).Equals(NewFranc(5)))
	})
	t.Run("Money currency", func(t *testing.T) {
		assert.Equal(t, "USD", NewDollar(1).Currency())
		assert.Equal(t, "CHF", NewFranc(1).Currency())
	})
	t.Run("Simple addition", func(t *testing.T) {
		five := NewDollar(5)
		sum := five.Plus(five)
		bank := NewBank()
		reduced := bank.Reduce(sum, "USD")
		assert.Equal(t, NewDollar(10), reduced)
	})
	t.Run("Plus returns sum", func(t *testing.T) {
		five := NewDollar(10)
		result := five.Plus(five)
		s, ok := result.(Sum)
		assert.True(t, ok)
		assert.Equal(t, five, s.augend)
		assert.Equal(t, five, s.addend)
	})
	t.Run("Reduce sum", func(t *testing.T) {
		sum := Bank{}.Reduce(
			NewSum(
				NewDollar(3), NewDollar(4),
			),
			"USD",
		)
		assert.Equal(t, NewDollar(7), sum)
	})
	t.Run("Reduce money different currency", func(t *testing.T) {
		bank := NewBank()
		rate := 2

		bank.AddRate("CHF", "USD", rate)
		assert.Equal(t, bank.Rate("CHF", "USD"), rate)

		result := bank.Reduce(NewFranc(2), "USD")
		assert.Equal(t, NewDollar(1), result)
	})
	t.Run("Mixed addition", func(t *testing.T) {
		fiveBucks := NewDollar(5)
		tenFrancs := NewFranc(10)
		bank := NewBank()
		bank.AddRate("CHF", "USD", 2)
		result := bank.Reduce(fiveBucks.Plus(tenFrancs), "USD")
		assert.Equal(t, NewDollar(10), result)
	})
}
