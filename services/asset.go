package services

import (
	"time"

	"github.com/Siddheshk02/portfolio-management/models"
	"github.com/Siddheshk02/portfolio-management/repositories"
)

type AssetService interface {
	AddAsset(portfolioID uint, name string, value float64, date time.Time) error
	GetAssetByID(id uint) (*models.Asset, error)
	UpdateAsset(id uint, name string, value float64, date time.Time) error
	DeleteAsset(id uint) error
}

type assetService struct {
	assetRepo repositories.AssetRepository
}

func NewAssetService(assetRepo repositories.AssetRepository) AssetService {
	return &assetService{assetRepo: assetRepo}
}

func (s *assetService) AddAsset(portfolioID uint, name string, value float64, date time.Time) error {
	asset := &models.Asset{
		PortfolioID: portfolioID,
		Name:        name,
		Value:       value,
		CreatedAt:   date,
	}
	return s.assetRepo.AddAsset(asset)
}

func (s *assetService) GetAssetByID(id uint) (*models.Asset, error) {
	return s.assetRepo.GetAssetByID(id)
}

func (s *assetService) UpdateAsset(id uint, name string, value float64, date time.Time) error {
	asset, err := s.assetRepo.GetAssetByID(id)
	if err != nil {
		return err
	}
	asset.Name = name
	asset.Value = value
	asset.CreatedAt = date
	return s.assetRepo.UpdateAsset(asset)
}

func (s *assetService) DeleteAsset(id uint) error {
	return s.assetRepo.DeleteAsset(id)
}
