package repository

import (
	"context"
	"errors"

	"github.com/hassan-alidoost/oms-project/internal/domain"
	"github.com/hassan-alidoost/oms-project/internal/infra/postgres/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}
 
func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}
 
func (r *ProductRepository) FindByID(ctx context.Context, id uint64) (*domain.Product, error) {
	var m models.Product
	result := r.db.WithContext(ctx).First(&m, "id = ?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrProductNotFound
		}
		return nil, err
	}
	return m.ToDomain(), nil
}
 
func (r *ProductRepository) FindAll(ctx context.Context) ([]*domain.Product, error) {
	var models []models.Product
	if err := r.db.WithContext(ctx).Order("created_at ASC").Find(&models).Error; err != nil {
		return nil, err
	}
	products := make([]*domain.Product, len(models))
	for i, m := range models {
		m := m
		products[i] = m.ToDomain()
	}
	return products, nil
}
 
func (r *ProductRepository) Save(ctx context.Context, p *domain.Product) error {
	m := models.ToProductModel(p)
	return r.db.WithContext(ctx).Create(m).Error
}
 
func (r *ProductRepository) Update(ctx context.Context, p *domain.Product) error {
	m := models.ToProductModel(p)
	result := r.db.WithContext(ctx).
		Model(m).
		Where("id = ?", p.ID).
		Updates(map[string]any{
			"name":       m.Name,
			"sku":        m.SKU,
			"price":      m.Price,
			"stock":      m.Stock,
			"updated_at": m.UpdatedAt,
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrProductNotFound
	}
	return nil
}
