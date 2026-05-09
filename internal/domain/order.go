package domain

type OrderStatus uint8

type Order struct {
	BaseEntity
	
	UserID     EntityId
	Status     OrderStatus
	TotalPrice Price
	Items      []OrderItem
}

const (
	Pending OrderStatus = iota
	Paid
	Shipped
	Cancelled
)

type OrderItem struct {
	BaseEntity

	Quantity  uint8
	Price     Price
}
