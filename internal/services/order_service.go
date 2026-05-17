package services

import (
	"context"
	"fmt"

	"github.com/hassan-alidoost/oms-project/internal/domain"
	"github.com/hassan-alidoost/oms-project/internal/handlers/http/order"
	"github.com/hassan-alidoost/oms-project/internal/ports"
)

type OrderService struct {
	orderRepo   ports.OrderRepository
	productRepo ports.ProductRepository
}

func NewOrderService(orderRepo ports.OrderRepository, productRepo ports.ProductRepository) *OrderService {
	return &OrderService{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

func (s *OrderService) Create(ctx context.Context, dto order.CreateDTO) (*order.ResponseDTO, error) {
	if len(dto.Items) == 0 {
		return nil, domain.ErrOrderEmpty
	}

	items := make([]domain.Item, 0, len(dto.Items))
	var reservations []stockReservation

	for _, itemDTO := range dto.Items {
		p, err := s.productRepo.FindByID(ctx, itemDTO.ProductID)
		if err != nil {
			s.rollbackStock(ctx, reservations)
			return nil, fmt.Errorf("App.Order.Create: product %s: %w", itemDTO.ProductID, err)
		}
		if err := p.DeductStock(itemDTO.Quantity); err != nil {
			s.rollbackStock(ctx, reservations)
			return nil, fmt.Errorf("App.Order.Create: product %s: %w", itemDTO.ProductID, err)
		}
		if err := s.productRepo.Update(ctx, p); err != nil {
			s.rollbackStock(ctx, reservations)
			return nil, fmt.Errorf("App.Order.Create: updating stock for product %s: %w", itemDTO.ProductID, err)
		}
		reservations = append(reservations, stockReservation{product: p, qty: itemDTO.Quantity})
		items = append(items, domain.Item{
			ProductID:   p.ID,
			ProductName: p.Name,
			SKU:         p.SKU,
			Quantity:    itemDTO.Quantity,
			UnitPrice:   p.Price,
		})
	}

	o, err := domain.NewOrder(items)
	if err != nil {
		s.rollbackStock(ctx, reservations)
		return nil, fmt.Errorf("App.Order.Create: %w", err)
	}
	if err := s.orderRepo.Save(ctx, o); err != nil {
		s.rollbackStock(ctx, reservations)
		return nil, fmt.Errorf("App.Order.Create: %w", err)
	}
	return toResponseDTO(o), nil
}

func (s *OrderService) GetByID(ctx context.Context, id uint64) (*order.ResponseDTO, error) {
	o, err := s.orderRepo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("App.Order.GetByID: %w", err)
	}
	return toResponseDTO(o), nil
}

func (s *OrderService) GetAll(ctx context.Context) ([]*order.ResponseDTO, error) {
	orders, err := s.orderRepo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("App.Order.GetAll: %w", err)
	}
	return toDTOList(orders), nil
}

func (s *OrderService) UpdateStatus(ctx context.Context, id uint64, dto order.UpdateStatusDTO) (*order.ResponseDTO, error) {
	o, err := s.orderRepo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("App.Order.UpdateStatus: %w", err)
	}
	if err := o.Transition(domain.OrderStatus(dto.Status)); err != nil {
		return nil, fmt.Errorf("App.Order.UpdateStatus: %w", err)
	}
	if err := s.orderRepo.Update(ctx, o); err != nil {
		return nil, fmt.Errorf("App.Order.UpdateStatus: %w", err)
	}
	return toResponseDTO(o), nil
}

type stockReservation struct {
	product *domain.Product
	qty     int
}

func (s *OrderService) rollbackStock(ctx context.Context, reservations []stockReservation) {
	for _, r := range reservations {
		r.product.RestoreStock(r.qty)
		_ = s.productRepo.Update(ctx, r.product)
	}
}

func toDTOList(orders []*domain.Order) []*order.ResponseDTO {
	dtos := make([]*order.ResponseDTO, len(orders))
	for i, o := range orders {
		dtos[i] = toResponseDTO(o)
	}
	return dtos
}

func toResponseDTO(o *domain.Order) *order.ResponseDTO {
	items := make([]order.ItemResponseDTO, len(o.Items))
	for i, item := range o.Items {
		items[i] = order.ItemResponseDTO{
			ProductID:   item.ProductID,
			ProductName: item.ProductName,
			SKU:         item.SKU,
			Quantity:    item.Quantity,
			UnitPrice:   item.UnitPrice,
			Subtotal:    item.Subtotal(),
		}
	}
	return &order.ResponseDTO{
		ID:          o.ID,
		Items:       items,
		Status:      string(o.Status),
		TotalAmount: o.TotalAmount,
		CreatedAt:   o.CreatedAt,
		UpdatedAt:   o.UpdatedAt,
	}
}
