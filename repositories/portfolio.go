package repositories

import (
	"github.com/Siddheshk02/portfolio-management/models"
	"github.com/jinzhu/gorm"
)

type PortfolioRepository interface {
	CreatePortfolio(portfolio *models.Portfolio) error
	GetPortfolioByID(id uint) (*models.Portfolio, error)
	UpdatePortfolio(portfolio *models.Portfolio) error
	DeletePortfolio(id uint) error
}

type portfolioRepository struct {
	db *gorm.DB
}

func NewPortfolioRepository(db *gorm.DB) PortfolioRepository {
	return &portfolioRepository{db: db}
}

func (r *portfolioRepository) CreatePortfolio(portfolio *models.Portfolio) error {
	return r.db.Create(portfolio).Error
}

func (r *portfolioRepository) GetPortfolioByID(id uint) (*models.Portfolio, error) {
	var portfolio models.Portfolio
	err := r.db.Preload("Assets").First(&portfolio, id).Error

	return &portfolio, err
}

func (r *portfolioRepository) UpdatePortfolio(portfolio *models.Portfolio) error {
	return r.db.Save(portfolio).Error
}

func (r *portfolioRepository) DeletePortfolio(id uint) error {
	return r.db.Delete(&models.Portfolio{}, id).Error
}
