package models

import (
	"time"
)

type Asset struct {
	ID          uint      `gorm:"primary_key"`
	PortfolioID uint      `gorm:"not null"`
	Name        string    `gorm:"not null"`
	Value       float64   `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null"`
}
