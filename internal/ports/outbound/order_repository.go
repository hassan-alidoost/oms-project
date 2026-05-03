package outbound

import (
	"context"

	"github.com/hassan-alidoost/oms-project/internal/domain/entities"
)

type OrderRepository interface {
    Create(ctx context.Context, order *entities.Order) error
	FindByID(ctx context.Context, id entities.EntityId) (*entities.Order, error)
}