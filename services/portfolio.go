package services

import (
	"github.com/Siddheshk02/portfolio-management/models"
	"github.com/Siddheshk02/portfolio-management/repositories"
)

type PortfolioService interface {
	CreatePortfolio(userID uint, name string) error
	GetPortfolioByID(id uint) (*models.Portfolio, error)
	UpdatePortfolio(id uint, name string) error
	DeletePortfolio(id uint) error
	CalculateTotalValue(portfolioID uint) (float64, error)
	CalculateAverageReturn(portfolioID uint) (float64, error)
}

type portfolioService struct {
	portfolioRepo repositories.PortfolioRepository
	assetRepo     repositories.AssetRepository
}

func NewPortfolioService(portfolioRepo repositories.PortfolioRepository, assetRepo repositories.AssetRepository) PortfolioService {
	return &portfolioService{
		portfolioRepo: portfolioRepo,
		assetRepo:     assetRepo,
	}
}

func (s *portfolioService) CreatePortfolio(userID uint, name string) error {
	portfolio := &models.Portfolio{
		UserID: userID,
		Name:   name,
	}

	return s.portfolioRepo.CreatePortfolio(portfolio)
}

func (s *portfolioService) GetPortfolioByID(id uint) (*models.Portfolio, error) {
	return s.portfolioRepo.GetPortfolioByID(id)
}

func (s *portfolioService) UpdatePortfolio(id uint, name string) error {
	portfolio, err := s.portfolioRepo.GetPortfolioByID(id)
	if err != nil {
		return err
	}

	portfolio.Name = name
	return s.portfolioRepo.UpdatePortfolio(portfolio)
}

func (s *portfolioService) DeletePortfolio(id uint) error {
	return s.portfolioRepo.DeletePortfolio(id)
}

func (s *portfolioService) CalculateTotalValue(portfolioID uint) (float64, error) {
	assets, err := s.assetRepo.GetAssetsByPortfolioID(portfolioID)
	if err != nil {
		return 0, err
	}

	totalValue := 0.0
	for _, asset := range assets {
		totalValue += asset.Value
	}

	return totalValue, nil
}

func (s *portfolioService) CalculateAverageReturn(portfolioID uint) (float64, error) {
	assets, err := s.assetRepo.GetAssetsByPortfolioID(portfolioID)
	if err != nil {
		return 0, err
	}

	var totalReturn float64
	var count int

	for _, asset := range assets {
		totalReturn += asset.Value
		count++
	}

	if count == 0 {
		return 0, nil
	}

	averageReturn := totalReturn / float64(count)
	return averageReturn, nil
}
