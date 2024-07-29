package repositories

import (
	"github.com/Siddheshk02/portfolio-management/models"
	"github.com/jinzhu/gorm"
)

type AssetRepository interface {
	AddAsset(asset *models.Asset) error
	GetAssetByID(id uint) (*models.Asset, error)
	UpdateAsset(asset *models.Asset) error
	DeleteAsset(id uint) error
	GetAssetsByPortfolioID(portfolioID uint) ([]*models.Asset, error)
}

type assetRepository struct {
	db *gorm.DB
}

func NewAssetRepository(db *gorm.DB) AssetRepository {
	return &assetRepository{db: db}
}

func (r *assetRepository) AddAsset(asset *models.Asset) error {
	return r.db.Create(asset).Error
}

func (r *assetRepository) GetAssetByID(id uint) (*models.Asset, error) {
	var asset models.Asset
	err := r.db.First(&asset, id).Error
	return &asset, err
}

func (r *assetRepository) UpdateAsset(asset *models.Asset) error {
	return r.db.Save(asset).Error
}

func (r *assetRepository) DeleteAsset(id uint) error {
	return r.db.Delete(&models.Asset{}, id).Error
}

func (r *assetRepository) GetAssetsByPortfolioID(portfolioID uint) ([]*models.Asset, error) {
	var assets []*models.Asset
	if err := r.db.Where("portfolio_id = ?", portfolioID).Find(&assets).Error; err != nil {
		return nil, err
	}
	return assets, nil
}
