package models

import (
	"time"
)

type User struct {
	ID         uint   `gorm:"primary_key"`
	Username   string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Portfolios []Portfolio `json:"portfolios"`
}
