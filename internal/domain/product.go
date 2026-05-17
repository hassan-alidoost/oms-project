package domain

import (
	"errors"
	"time"
)

var (
	ErrProductNotFound    = errors.New("product not found")
	ErrInvalidProduct     = errors.New("invalid product")
	ErrInsufficientStock  = errors.New("insufficient stock")
	ErrProductNameEmpty   = errors.New("product name cannot be empty")
	ErrProductPriceInvalid = errors.New("product price must be greater than zero")
)

type Product struct {
	ID        uint64
	Name      string
	SKU       string
	Price     float64
	Stock     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewProduct(id uint64, name, sku string, price float64, stock int) (*Product, error) {
	if name == "" {
		return nil, ErrProductNameEmpty
	}
	if price <= 0 {
		return nil, ErrProductPriceInvalid
	}
	now := time.Now().UTC()
	return &Product{
		ID:        id,
		Name:      name,
		SKU:       sku,
		Price:     price,
		Stock:     stock,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (p *Product) DeductStock(qty int) error {
	if p.Stock < qty {
		return ErrInsufficientStock
	}
	p.Stock -= qty
	p.UpdatedAt = time.Now().UTC()
	return nil
}

func (p *Product) RestoreStock(qty int) {
	p.Stock += qty
	p.UpdatedAt = time.Now().UTC()
}

func (p *Product) UpdateDetails(name, sku string, price float64, stock int) error {
	if name == "" {
		return ErrProductNameEmpty
	}
	if price <= 0 {
		return ErrProductPriceInvalid
	}
	p.Name = name
	p.SKU = sku
	p.Price = price
	p.Stock = stock
	p.UpdatedAt = time.Now().UTC()
	return nil
}