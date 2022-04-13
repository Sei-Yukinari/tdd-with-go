package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiCurrencyMoney(t *testing.T) {
	t.Run("Can be multiplied by dollars", func(t *testing.T) {
		five := NewDollar(5)
		assert.Equal(t, NewDollar(10), five.times(2))
		assert.Equal(t, NewDollar(15), five.times(3))
	})
	t.Run("Same amount of dollars is equivalent", func(t *testing.T) {
		assert.True(t, NewDollar(5).equals(NewDollar(5)))
		assert.False(t, NewDollar(5).equals(NewDollar(6)))
	})
}
