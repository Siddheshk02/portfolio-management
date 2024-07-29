package services

import (
	"testing"

	"github.com/Siddheshk02/portfolio-management/models"
	"github.com/Siddheshk02/portfolio-management/repositories"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestAddPortfolio(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	db.AutoMigrate(&models.User{}, &models.Portfolio{}, &models.Asset{})

	userRepo := repositories.NewUserRepository(db)
	userService := NewUserService(userRepo)

	err = userService.RegisterUser("Test User", "password")
	assert.NoError(t, err)

	user, err := userRepo.GetUserByUsername("Test User")
	assert.NoError(t, err)

	password := "password"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	userRepo.CreateUser(&models.User{Username: "Test User", Password: string(hashedPassword)})

	_, err = userService.AuthenticateUser("Test User", password)

	portfolioRepo := repositories.NewPortfolioRepository(db)
	assetRepo := repositories.NewAssetRepository(db)
	portfolioService := NewPortfolioService(portfolioRepo, assetRepo)

	err = portfolioService.CreatePortfolio(user.ID, "Test Portfolio")
	assert.NoError(t, err)

	portfolio, err := portfolioRepo.GetPortfolioByID(1)
	assert.NoError(t, err)
	assert.Equal(t, "Test Portfolio", portfolio.Name)
	assert.Equal(t, user.ID, portfolio.UserID)

	db.DropTable("portfolios")
	db.DropTable("users")
	db.DropTable("assets")
}

func TestGetTotalValue(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	db.AutoMigrate(&models.User{}, &models.Portfolio{}, &models.Asset{})

	portfolioRepo := repositories.NewPortfolioRepository(db)
	assetRepo := repositories.NewAssetRepository(db)
	portfolioService := NewPortfolioService(portfolioRepo, assetRepo)

	// user := &models.User{Username: "Test User", Password: "password"}
	// db.Create(user)

	portfolio := &models.Portfolio{UserID: 1, Name: "Test Portfolio"}
	db.Create(portfolio)

	asset1 := &models.Asset{PortfolioID: portfolio.ID, Name: "Asset 1", Value: 100}
	asset2 := &models.Asset{PortfolioID: portfolio.ID, Name: "Asset 2", Value: 200}
	db.Create(asset1)
	db.Create(asset2)

	totalValue, err := portfolioService.CalculateTotalValue(portfolio.ID)
	assert.NoError(t, err)
	assert.Equal(t, 300.0, totalValue)
}

func TestCalculateAverageReturn(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	db.AutoMigrate(&models.User{}, &models.Portfolio{}, &models.Asset{})

	portfolioRepo := repositories.NewPortfolioRepository(db)
	assetRepo := repositories.NewAssetRepository(db)
	portfolioService := NewPortfolioService(portfolioRepo, assetRepo)

	// user := &models.User{Username: "Test User", Password: "password"}
	// db.Create(user)

	portfolio := &models.Portfolio{UserID: 1, Name: "Test Portfolio"}
	db.Create(portfolio)

	asset1 := &models.Asset{PortfolioID: portfolio.ID, Name: "Asset 1", Value: 100}
	asset2 := &models.Asset{PortfolioID: portfolio.ID, Name: "Asset 2", Value: 200}
	db.Create(asset1)
	db.Create(asset2)

	averageReturn, err := portfolioService.CalculateAverageReturn(portfolio.ID)
	assert.NoError(t, err)
	assert.Equal(t, 150.0, averageReturn)
}
