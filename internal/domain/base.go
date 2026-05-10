package domain

import "time"

type Price uint64
type ID uint64

type Base struct {
	ID        ID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
