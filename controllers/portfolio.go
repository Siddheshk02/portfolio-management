package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Siddheshk02/portfolio-management/models"
	"github.com/Siddheshk02/portfolio-management/services"
	"github.com/gorilla/mux"
)

type PortfolioController struct {
	portfolioService services.PortfolioService
}

func NewPortfolioController(portfolioService services.PortfolioService) *PortfolioController {
	return &PortfolioController{portfolioService: portfolioService}
}

func (c *PortfolioController) CreatePortfolio(w http.ResponseWriter, r *http.Request) {
	var portfolio models.Portfolio

	err := json.NewDecoder(r.Body).Decode(&portfolio)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("userID").(uint)
	err = c.portfolioService.CreatePortfolio(userID, portfolio.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *PortfolioController) GetPortfolio(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid portfolio ID", http.StatusBadRequest)
		return
	}
	portfolio, err := c.portfolioService.GetPortfolioByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(portfolio)
}

func (c *PortfolioController) UpdatePortfolio(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid portfolio ID", http.StatusBadRequest)
		return
	}
	var portfolio models.Portfolio
	if err := json.NewDecoder(r.Body).Decode(&portfolio); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.portfolioService.UpdatePortfolio(uint(id), portfolio.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *PortfolioController) DeletePortfolio(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid portfolio ID", http.StatusBadRequest)
		return
	}
	err = c.portfolioService.DeletePortfolio(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c *PortfolioController) CalculateTotalValue(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid portfolio ID", http.StatusBadRequest)
		return
	}
	totalValue, err := c.portfolioService.CalculateTotalValue(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]float64{"total_value": totalValue})
}

func (c *PortfolioController) CalculateAverageReturn(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid portfolio ID", http.StatusBadRequest)
		return
	}

	averageReturn, err := c.portfolioService.CalculateAverageReturn(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]float64{"average_return": averageReturn})
}
