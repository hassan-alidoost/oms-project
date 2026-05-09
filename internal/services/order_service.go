package services

import (
	"context"
	"errors"
	"time"

	"github.com/hassan-alidoost/oms-project/internal/domain"
	"github.com/hassan-alidoost/oms-project/internal/ports"
)

type OrderService struct {
	repo ports.OrderRepository
}

func NewOrderService(repo ports.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) Create(ctx context.Context, userID domain.EntityId, price domain.Price) (*domain.Order, error) {
    newOrder := &domain.Order{
        UserID: userID,
        TotalPrice: price,
        Status:     domain.Pending,
		BaseEntity: domain.BaseEntity{
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
    }

    if err := s.repo.Create(ctx, newOrder); err != nil {
        return nil, err
    }

    return newOrder, nil
}

func (s *OrderService) FindByID(ctx context.Context, id domain.EntityId) (*domain.Order, error) {
	order, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errors.New("order not found")
	}
	return order, nil
}