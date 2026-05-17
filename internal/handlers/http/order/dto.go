package order

import (
	"time"

	"github.com/hassan-alidoost/oms-project/internal/domain"
)

type CreateItemDTO struct {
	ProductID uint64 `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type CreateDTO struct {
	Items []CreateItemDTO `json:"items"`
}

type UpdateStatusDTO struct {
	Status domain.OrderStatus `json:"status"`
}

type ItemResponseDTO struct {
	ProductID   uint64  `json:"product_id"`
	ProductName string  `json:"product_name"`
	SKU         string  `json:"sku"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	Subtotal    float64 `json:"subtotal"`
}

type ResponseDTO struct {
	ID          uint64            `json:"id"`
	Items       []ItemResponseDTO `json:"items"`
	Status      string            `json:"status"`
	TotalAmount float64           `json:"total_amount"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}
