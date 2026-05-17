package models

import (
	"time"

	"github.com/hassan-alidoost/oms-project/internal/domain"
)

type Product struct {
	ID        uint64    `gorm:"column:id;primaryKey"`
	Name      string    `gorm:"column:name;not null"`
	SKU       string    `gorm:"column:sku;not null;default:''"`
	Price     float64   `gorm:"column:price;not null"`
	Stock     int       `gorm:"column:stock;not null;default:0"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
 
func (Product) TableName() string { return "products" }
 
func ToProductModel(p *domain.Product) *Product {
	return &Product{
		ID:        p.ID,
		Name:      p.Name,
		SKU:       p.SKU,
		Price:     p.Price,
		Stock:     p.Stock,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
 
func (m *Product) ToDomain() *domain.Product {
	return &domain.Product{
		ID:        m.ID,
		Name:      m.Name,
		SKU:       m.SKU,
		Price:     m.Price,
		Stock:     m.Stock,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
