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
}
