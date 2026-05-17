package ports

import (
	"context"

	"github.com/hassan-alidoost/oms-project/internal/domain"
	"github.com/hassan-alidoost/oms-project/internal/handlers/http/order"
)
 
type OrderRepository interface {
	FindByID(ctx context.Context, id uint64) (*domain.Order, error)
	FindAll(ctx context.Context) ([]*domain.Order, error)
	Save(ctx context.Context, order *domain.Order) error
	Update(ctx context.Context, order *domain.Order) error
}
 
type OrderService interface {
	GetByID(ctx context.Context, id uint64) (*order.ResponseDTO, error)
	GetAll(ctx context.Context) ([]*order.ResponseDTO, error)
	Create(ctx context.Context, dto order.CreateDTO) (*order.ResponseDTO, error)
	UpdateStatus(ctx context.Context, id uint64, dto order.UpdateStatusDTO) (*order.ResponseDTO, error)
}
