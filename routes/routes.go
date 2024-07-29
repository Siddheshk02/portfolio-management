package routes

import (
	"github.com/Siddheshk02/portfolio-management/controllers"
	"github.com/Siddheshk02/portfolio-management/middlewares"
	"github.com/gorilla/mux"
)

func SetupRoutes(userController *controllers.UserController, portfolioController *controllers.PortfolioController, assetController *controllers.AssetController) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/register", userController.RegisterUser).Methods("POST")
	router.HandleFunc("/login", userController.Login).Methods("POST")

	protected := router.PathPrefix("/").Subrouter()
	protected.Use(middlewares.JWTAuth)
	protected.HandleFunc("/portfolios", portfolioController.CreatePortfolio).Methods("POST")
	protected.HandleFunc("/portfolios/{id}", portfolioController.GetPortfolio).Methods("GET")
	protected.HandleFunc("/portfolios/{id}", portfolioController.UpdatePortfolio).Methods("PUT")
	protected.HandleFunc("/portfolios/{id}", portfolioController.DeletePortfolio).Methods("DELETE")
	protected.HandleFunc("/portfolios/{id}/totalvalue", portfolioController.CalculateTotalValue).Methods("GET")
	protected.HandleFunc("/portfolios/{id}/averagereturn", portfolioController.CalculateAverageReturn).Methods("GET")

	protected.HandleFunc("/portfolios/{id}/assets", assetController.AddAsset).Methods("POST")
	protected.HandleFunc("/portfolios/{id}/assets/{asset_id}", assetController.GetAsset).Methods("GET")
	protected.HandleFunc("/portfolios/{id}/assets/{asset_id}", assetController.UpdateAsset).Methods("PUT")
	protected.HandleFunc("/portfolios/{id}/assets/{asset_id}", assetController.DeleteAsset).Methods("DELETE")

	return router
}
