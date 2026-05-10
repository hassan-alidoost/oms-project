package models

import (
	"github.com/hassan-alidoost/oms-project/internal/domain"
)

type Order struct {
	BaseModel
	Status     uint8        `gorm:"type:type:smallint;not null"`
	TotalPrice domain.Price `gorm:"type:bigint;not null"`

	Items []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	BaseModel
	OrderID   ID
	ProductID ID

	Product Product `gorm:"foreignKey:ProductID"`

	Quantity uint8
	Price    domain.Price
}

func ToModel(dOrder *domain.Order) *Order {
	mOrder := &Order{
		Status:     uint8(dOrder.Status),
		TotalPrice: dOrder.TotalPrice,
	}

	if dOrder.ID != 0 {
		mOrder.ID = ID(dOrder.ID)
	}

	for _, item := range dOrder.Items {
		mItem := OrderItem{
			ProductID: ID(item.ProductID),
			Quantity:  item.Quantity,
			Price:     item.Price,
		}
		if item.ID != 0 {
			mItem.ID = ID(item.ID)
		}
		mOrder.Items = append(mOrder.Items, mItem)
	}

	return mOrder
}

func ToDomain(mOrder *Order) *domain.Order {
	dOrder := &domain.Order{
		Base: domain.Base{
			ID:        domain.ID(mOrder.ID),
			CreatedAt: mOrder.CreatedAt,
			UpdatedAt: mOrder.UpdatedAt,
		},
		Status:     domain.OrderStatus(mOrder.Status),
		TotalPrice: mOrder.TotalPrice,
	}

	if mOrder.DeletedAt.Valid {
		dOrder.Base.DeletedAt = &mOrder.DeletedAt.Time
	}

	for _, item := range mOrder.Items {
		dItem := domain.OrderItem{
			Base: domain.Base{
				ID:        domain.ID(item.ID),
				CreatedAt: item.CreatedAt,
				UpdatedAt: item.UpdatedAt,
			},
			ProductID: domain.ID(item.ProductID),
			Quantity:  item.Quantity,
			Price:     item.Price,
		}

		if item.Product.ID != 0 {
			dItem.Product = &domain.Product{
				Base: domain.Base{
					ID:        domain.ID(item.Product.ID),
					CreatedAt: item.Product.CreatedAt,
					UpdatedAt: item.Product.UpdatedAt,
				},
				Name:  item.Product.Name,
				Price: domain.Price(item.Product.Price),
				Stock: item.Product.Stock,
			}
		}

		dOrder.Items = append(dOrder.Items, dItem)
	}

	return dOrder
}
