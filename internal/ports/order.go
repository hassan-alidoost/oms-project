package ports

import (
	"context"

	"github.com/hassan-alidoost/oms-project/internal/domain"
)

type OrderRepository interface {
    Create(ctx context.Context, order *domain.Order) error
	FindByID(ctx context.Context, id domain.EntityId) (*domain.Order, error)
}

type OrderService interface {
	Create(ctx context.Context, order *domain.Order) error
	FindByID(ctx context.Context, id domain.EntityId) (*domain.Order, error)
}