package services

import (
	"context"
	"errors"
	"time"

	"github.com/hassan-alidoost/oms-project/internal/domain/entities"
	"github.com/hassan-alidoost/oms-project/internal/ports/outbound"
)

type OrderService struct {
	repo outbound.OrderRepository
}

func NewOrderService(repo outbound.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) PlaceOrder(ctx context.Context, userID entities.EntityId, price entities.Price) (*entities.Order, error) {
    newOrder := &entities.Order{
        UserID: userID,
        TotalPrice: price,
        Status:     entities.Pending,
		BaseEntity: entities.BaseEntity{
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
    }

    if err := s.repo.Create(ctx, newOrder); err != nil {
        return nil, err
    }

    return newOrder, nil
}

func (s *OrderService) FindByID(ctx context.Context, id entities.EntityId) (*entities.Order, error) {
	order, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errors.New("order not found")
	}
	return order, nil
}