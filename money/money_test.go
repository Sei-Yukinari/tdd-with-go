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
}
