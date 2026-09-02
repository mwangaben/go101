package enums

import "slices"

type OrderStatus int

const (
	OrderCreated OrderStatus = iota
	OrderPending
	OrderConfirmed
	OrderProcessing
	OrderShipped
	OrderDelivered
	OrderCancelled
	OrderRefunded
)

func (s OrderStatus) String() string {
	return [...]string{"Created", "Pending", "Confirmed", "Processing", "Shipped", "Delivered", "Cancelled", "Refunded"}[s]
}

func (s OrderStatus) CanTransition(to OrderStatus) bool {
	transitions := map[OrderStatus][]OrderStatus{
		OrderCreated:    {OrderPending, OrderCancelled},
		OrderPending:    {OrderConfirmed, OrderCancelled},
		OrderConfirmed:  {OrderProcessing, OrderCancelled},
		OrderProcessing: {OrderShipped, OrderCancelled},
		OrderShipped:    {OrderDelivered},
		OrderDelivered:  {OrderRefunded},
		OrderCancelled:  {},
		OrderRefunded:   {},
	}

	validTransitions, ok := transitions[s]
	if !ok {
		return false
	}

	return slices.Contains(validTransitions, to)
}

func (s OrderStatus) IsTerminal() bool {
	return s == OrderCancelled || s == OrderRefunded || s == OrderDelivered
}

func (s OrderStatus) NextStates() []OrderStatus {
	transitions := map[OrderStatus][]OrderStatus{
		OrderCreated:    {OrderPending, OrderCancelled},
		OrderPending:    {OrderConfirmed, OrderCancelled},
		OrderConfirmed:  {OrderProcessing, OrderCancelled},
		OrderProcessing: {OrderShipped, OrderCancelled},
		OrderShipped:    {OrderDelivered},
		OrderDelivered:  {OrderRefunded},
		OrderCancelled:  {},
		OrderRefunded:   {},
	}

	return transitions[s]
}
