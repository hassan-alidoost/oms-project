package domain

import "time"

type OrderStatus uint8

type Order struct {
	Base

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
	Base
	ProductId ID
	Quantity uint8
	Price    Price
}

func NewOrderItem(productID ID, qty uint8, price Price) OrderItem {
    return OrderItem{
        Base: Base{
            CreatedAt: time.Now(),
            UpdatedAt: time.Now(),
        },
        ProductId: productID,
        Quantity:  qty,
        Price:     price,
    }
}

func NewOrder(items []OrderItem) *Order {
    order := &Order{
        Base: Base{
            CreatedAt: time.Now(),
            UpdatedAt: time.Now(),
        },
        Status:     Pending,
        Items:      items,
    }
    
    order.CalculateTotal()
    return order
}

func (o *Order) CalculateTotal() {
	var total Price
	for _, item := range o.Items {
		total += item.Price * Price(item.Quantity)
	}
	o.TotalPrice = total
}
