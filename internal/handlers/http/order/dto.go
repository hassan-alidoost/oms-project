package order

import "github.com/hassan-alidoost/oms-project/internal/domain"

type CreateOrderRequest struct {
	Items []OrderItemRequest `json:"items" validate:"required,min=1,dive"`
}

type OrderItemRequest struct {
	ProductID domain.ID    `json:"product_id" validate:"required"`
	Quantity  uint8        `json:"quantity" validate:"required,gt=0"`
	Price     domain.Price `json:"price" validate:"required,gt=0"`
}
