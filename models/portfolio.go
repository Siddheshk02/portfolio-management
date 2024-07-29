package models

import (
	"time"
)

type Portfolio struct {
	ID        uint    `gorm:"primary_key"`
	UserID    uint    `gorm:"not null"`
	Name      string  `gorm:"not null"`
	Assets    []Asset `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
