package enums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderStatus(t *testing.T) {

	t.Run("it test the valid transition", func(t *testing.T) {

		status := OrderCreated

		assert.Equal(t, status, OrderCreated)
		assert.Equal(t, "Created", status.String())
		assert.Equal(t, false, status.IsTerminal())
		assert.Equal(t, []OrderStatus{OrderPending, OrderCancelled}, status.NextStates())
		assert.Equal(t, true, status.CanTransition(OrderCancelled))
		assert.Equal(t, true, status.CanTransition(OrderPending))
		assert.Equal(t, false, status.CanTransition(OrderRefunded))

		status = OrderCancelled
		assert.Equal(t, status, OrderCancelled)
		assert.Equal(t, false, status.CanTransition(OrderRefunded))
		assert.Equal(t, false, status.CanTransition(OrderPending))
		assert.Equal(t, []OrderStatus{}, status.NextStates())

	})
}
