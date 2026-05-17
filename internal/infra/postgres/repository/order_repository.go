package repository

import (
	"context"
	"errors"

	"github.com/hassan-alidoost/oms-project/internal/domain"
	"github.com/hassan-alidoost/oms-project/internal/infra/postgres/models"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}
 
func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}
 
func (r *OrderRepository) FindByID(ctx context.Context, id uint64) (*domain.Order, error) {
	var m models.Order
	result := r.db.WithContext(ctx).
		Preload("Items").
		First(&m, "id = ?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrOrderNotFound
		}
		return nil, err
	}
	return m.ToDomain(), nil
}
 
func (r *OrderRepository) FindAll(ctx context.Context) ([]*domain.Order, error) {
	var models []models.Order
	if err := r.db.WithContext(ctx).
		Preload("Items").
		Order("created_at ASC").
		Find(&models).Error; err != nil {
		return nil, err
	}
	return ToOrderDomainList(models), nil
}
 
func (r *OrderRepository) Save(ctx context.Context, o *domain.Order) error {
	m := models.ToOrderModel(o)
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit("Items").Create(m).Error; err != nil {
			return err
		}

		if len(m.Items) > 0 {
			if err := tx.Create(&m.Items).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
 
func (r *OrderRepository) Update(ctx context.Context, o *domain.Order) error {
	result := r.db.WithContext(ctx).
		Model(&models.Order{}).
		Where("id = ?", o.ID).
		Updates(map[string]any{
			"status":       models.OrderStatus(o.Status),
			"total_amount": o.TotalAmount,
			"updated_at":   o.UpdatedAt,
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrOrderNotFound
	}
	return nil
}
 
 
func ToOrderDomainList(models []models.Order) []*domain.Order {
	orders := make([]*domain.Order, len(models))
	for i, m := range models {
		m := m 
		orders[i] = m.ToDomain()
	}
	return orders
}
