package services

import (
	"context"
	"errors"

	"github.com/hassan-alidoost/oms-project/internal/domain"
	"github.com/hassan-alidoost/oms-project/internal/ports"
)

type OrderService struct {
	repo ports.OrderRepository
}

func NewOrderService(repo ports.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) Create(ctx context.Context, items []domain.OrderItem) error {
	order := domain.NewOrder(items)

	if err := s.repo.Create(ctx, order); err != nil {
		return err
	}

	return nil
}

func (s *OrderService) FindByID(ctx context.Context, id domain.ID) (*domain.Order, error) {
	order, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errors.New("order not found")
	}
	return order, nil
}
