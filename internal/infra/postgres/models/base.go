package models

import (
	"time"

	"gorm.io/gorm"
)

type Price uint64
type EntityId uint64

type BaseModel struct {
	ID        EntityId       `gorm:"primarykey;autoIncrement"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
