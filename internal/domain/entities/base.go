package entities

import "time"

type Price uint64
type EntityId uint64

type BaseEntity struct {
	ID        EntityId
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
