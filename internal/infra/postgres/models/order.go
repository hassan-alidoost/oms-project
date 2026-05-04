package models

import "github.com/hassan-alidoost/oms-project/internal/domain/entities"

type Order struct {
	BaseModel

	UserID     EntityId `gorm:"index"`
	Status     uint8    `gorm:"type:type:smallint;not null"`
	TotalPrice Price    `gorm:"type:bigint;not null"`
}

func FromDomain(d *entities.Order) *Order {
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

func (m *Order) ToDomain() *entities.Order {
	return &entities.Order{
		UserID:     entities.EntityId(m.UserID),
		Status:     entities.OrderStatus(m.Status),
		TotalPrice: entities.Price(m.TotalPrice),
		BaseEntity: entities.BaseEntity{
			ID:        entities.EntityId(m.ID),
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		},
	}
}
