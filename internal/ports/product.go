package ports

import (
	"context"

	"github.com/hassan-alidoost/oms-project/internal/domain"
)

type ProductRepository interface {
	FindByID(ctx context.Context, id uint64) (*domain.Product, error)
	FindAll(ctx context.Context) ([]*domain.Product, error)
	Save(ctx context.Context, product *domain.Product) error
	Update(ctx context.Context, product *domain.Product) error
}
 
