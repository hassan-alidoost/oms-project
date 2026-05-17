package domain

import (
	"errors"
	"time"
)

var (
	ErrOrderNotFound         = errors.New("order not found")
	ErrOrderEmpty            = errors.New("order must contain at least one item")
	ErrInvalidStatus         = errors.New("invalid order status transition")
	ErrOrderAlreadyCancelled = errors.New("order is already cancelled")
	ErrOrderAlreadyShipped   = errors.New("cannot cancel a shipped order")
	ErrCustomerIDEmpty       = errors.New("customer ID cannot be empty")
)

type OrderStatus uint8

const (
	StatusPending OrderStatus = iota
	StatusConfirmed
	StatusShipped
	StatusDelivered
	StatusCancelled
)

var validTransitions = map[OrderStatus][]OrderStatus{
	StatusPending:   {StatusConfirmed, StatusCancelled},
	StatusConfirmed: {StatusShipped, StatusCancelled},
	StatusShipped:   {StatusDelivered},
	StatusDelivered: {},
	StatusCancelled: {},
}

type Item struct {
	ProductID   uint64
	ProductName string
	SKU         string
	Quantity    int
	UnitPrice   float64
}

func (i Item) Subtotal() float64 {
	return float64(i.Quantity) * i.UnitPrice
}

type Order struct {
	ID          uint64
	Items       []Item
	Status      OrderStatus
	TotalAmount float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewOrder(items []Item) (*Order, error) {
	if len(items) == 0 {
		return nil, ErrOrderEmpty
	}

	total := 0.0
	for _, item := range items {
		total += item.Subtotal()
	}

	now := time.Now().UTC()
	return &Order{
		Items:       items,
		Status:      StatusPending,
		TotalAmount: total,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (o *Order) Transition(next OrderStatus) error {
	allowed, ok := validTransitions[o.Status]
	if !ok {
		return ErrInvalidStatus
	}
	for _, s := range allowed {
		if s == next {
			o.Status = next
			o.UpdatedAt = time.Now().UTC()
			return nil
		}
	}
	return ErrInvalidStatus
}

func (o *Order) Cancel() error {
	if o.Status == StatusCancelled {
		return ErrOrderAlreadyCancelled
	}
	if o.Status == StatusShipped || o.Status == StatusDelivered {
		return ErrOrderAlreadyShipped
	}
	return o.Transition(StatusCancelled)
}

func (o *Order) IsCancelled() bool {
	return o.Status == StatusCancelled
}
