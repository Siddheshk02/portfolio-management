package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Siddheshk02/portfolio-management/config"
	"github.com/Siddheshk02/portfolio-management/controllers"
	"github.com/Siddheshk02/portfolio-management/models"
	"github.com/Siddheshk02/portfolio-management/repositories"
	"github.com/Siddheshk02/portfolio-management/routes"
	"github.com/Siddheshk02/portfolio-management/services"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	config.LoadConfig()
	user := os.Getenv("User")
	dbname := os.Getenv("DB")
	pass := os.Getenv("Password")

	db, err := gorm.Open("postgres", "host=localhost port=5432 user="+user+" dbname="+dbname+" password="+pass+" sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	db.AutoMigrate(&models.User{}, &models.Portfolio{}, &models.Asset{})

	userRepo := repositories.NewUserRepository(db)
	portfolioRepo := repositories.NewPortfolioRepository(db)
	assetRepo := repositories.NewAssetRepository(db)

	userService := services.NewUserService(userRepo)
	portfolioService := services.NewPortfolioService(portfolioRepo, assetRepo)
	assetService := services.NewAssetService(assetRepo)

	userController := controllers.NewUserController(userService)
	portfolioController := controllers.NewPortfolioController(portfolioService)
	assetController := controllers.NewAssetController(assetService)

	router := routes.SetupRoutes(userController, portfolioController, assetController)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
