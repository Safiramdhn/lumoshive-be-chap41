package models

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id" example:"1"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name" example:"John Doe"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email" example:"john.doe@example.com"`
	Password  string    `gorm:"type:varchar(255);not null" json:"password" swaggerignore:"true"` // Sensitive field
	Points    int       `gorm:"type:int;default:0" json:"points" example:"100"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at" example:"2024-12-03T12:00:00Z"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at" example:"2024-12-03T12:30:00Z"`

	// Relationships
	Redeems []Redeem `gorm:"foreignKey:UserID" json:"redeems,omitempty"`
	Usages  []Usage  `gorm:"foreignKey:UserID" json:"usages,omitempty"`
}
