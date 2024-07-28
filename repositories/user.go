package repositories

import {
	"github.com/jinzhu/gorm"
	"github.com/Siddheshk02/portfolio-management/models"
}

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByUsername(username string) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db : db}
}

func (r *userRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}