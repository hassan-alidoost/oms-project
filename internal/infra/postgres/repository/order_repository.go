package repository

import (
	"context"
	"errors"

	"github.com/hassan-alidoost/oms-project/internal/domain"
	"github.com/hassan-alidoost/oms-project/internal/infra/postgres/models"
	"github.com/hassan-alidoost/oms-project/internal/ports"
	"gorm.io/gorm"
)


type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) ports.OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) Create(ctx context.Context, order *domain.Order) error {
	model := models.FromDomain(order)

	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return err
	}

	return nil
}

func (r *orderRepository) FindByID(ctx context.Context, id domain.EntityId) (*domain.Order, error) {
	var model models.Order

	err := r.db.WithContext(ctx).First(&model, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return model.ToDomain(), nil
}