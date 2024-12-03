package models

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        int            `gorm:"primaryKey;autoIncrement" json:"id,omitempty" swagger:"ignore"`
	CreatedAt time.Time      `json:"-" gorm:"default:CURRENT_TIMESTAMP" swagger:"ignore"`
	UpdatedAt time.Time      `json:"-" gorm:"default:CURRENT_TIMESTAMP" swagger:"ignore"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index" swagger:"ignore"`
}
