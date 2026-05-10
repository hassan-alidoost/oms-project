package models

import (
	"time"

	"gorm.io/gorm"
)

type Price uint64
type ID uint64

type BaseModel struct {
	ID        ID             `gorm:"primarykey;autoIncrement"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
