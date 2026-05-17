package models

import (
	"time"

	"github.com/hassan-alidoost/oms-project/internal/domain"
)

type OrderStatus uint8

type OrderItem struct {
	ID          uint64  `gorm:"column:id;primaryKey;autoIncrement"`
	OrderID     uint64  `gorm:"column:order_id;not null;index"`
	ProductID   uint64  `gorm:"column:product_id;not null"`
	ProductName string  `gorm:"column:product_name;not null;default:''"`
	SKU         string  `gorm:"column:sku;not null;default:''"`
	Quantity    int     `gorm:"column:quantity;not null"`
	UnitPrice   float64 `gorm:"column:unit_price;not null"`
}

func (OrderItem) TableName() string { return "order_items" }

type Order struct {
	ID          uint64           `gorm:"column:id;primaryKey"`
	Status      OrderStatus      `gorm:"column:status;not null;default:'pending'"`
	TotalAmount float64          `gorm:"column:total_amount;not null;default:0"`
	Notes       string           `gorm:"column:notes;not null;default:''"`
	Items       []OrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time        `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time        `gorm:"column:updated_at;autoUpdateTime"`
}

func (Order) TableName() string { return "orders" }

func ToOrderModel(o *domain.Order) *Order {
	items := make([]OrderItem, len(o.Items))
	for i, item := range o.Items {
		items[i] = OrderItem{
			OrderID:     o.ID,
			ProductID:   item.ProductID,
			ProductName: item.ProductName,
			SKU:         item.SKU,
			Quantity:    item.Quantity,
			UnitPrice:   item.UnitPrice,
		}
	}
	return &Order{
		ID:          o.ID,
		Status:      OrderStatus(o.Status),
		TotalAmount: o.TotalAmount,
		Items:       items,
		CreatedAt:   o.CreatedAt,
		UpdatedAt:   o.UpdatedAt,
	}
}

func (m *Order) ToDomain() *domain.Order {
	items := make([]domain.Item, len(m.Items))
	for i, item := range m.Items {
		items[i] = domain.Item{
			ProductID:   item.ProductID,
			ProductName: item.ProductName,
			SKU:         item.SKU,
			Quantity:    item.Quantity,
			UnitPrice:   item.UnitPrice,
		}
	}
	return &domain.Order{
		ID:          m.ID,
		Status:      domain.OrderStatus(m.Status),
		TotalAmount: m.TotalAmount,
		Items:       items,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}
