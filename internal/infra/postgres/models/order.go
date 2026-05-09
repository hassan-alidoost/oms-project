package models

import "github.com/hassan-alidoost/oms-project/internal/domain"


type Order struct {
	BaseModel

	UserID     EntityId `gorm:"index"`
	Status     uint8    `gorm:"type:type:smallint;not null"`
	TotalPrice Price    `gorm:"type:bigint;not null"`
}

func FromDomain(d *domain.Order) *Order {
	return &Order{
		UserID:     EntityId(d.UserID),
		Status:     uint8(d.Status),
		TotalPrice: Price(d.TotalPrice),
		BaseModel: BaseModel{
			ID:        EntityId(d.ID),
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		},
	}
}

func (m *Order) ToDomain() *domain.Order {
	return &domain.Order{
		UserID:     domain.EntityId(m.UserID),
		Status:     domain.OrderStatus(m.Status),
		TotalPrice: domain.Price(m.TotalPrice),
		BaseEntity: domain.BaseEntity{
			ID:        domain.EntityId(m.ID),
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		},
	}
}
