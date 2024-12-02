package models

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        int            `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	CreatedAt time.Time      `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
