package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Siddheshk02/portfolio-management/models"
	"github.com/Siddheshk02/portfolio-management/services"
	"github.com/gorilla/mux"
)

type AssetController struct {
	assetService services.AssetService
}

func NewAssetController(assetService services.AssetService) *AssetController {
	return &AssetController{assetService: assetService}
}

func (c *AssetController) AddAsset(w http.ResponseWriter, r *http.Request) {
	portfolioIDStr := mux.Vars(r)["id"]
	portfolioID, err := strconv.Atoi(portfolioIDStr)
	if err != nil {
		http.Error(w, "Invalid portfolio ID", http.StatusBadRequest)
		return
	}
	var asset models.Asset
	if err := json.NewDecoder(r.Body).Decode(&asset); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	asset.PortfolioID = uint(portfolioID)
	asset.CreatedAt = time.Now()
	err = c.assetService.AddAsset(asset.PortfolioID, asset.Name, asset.Value, asset.CreatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *AssetController) GetAsset(w http.ResponseWriter, r *http.Request) {
	portfolioIDStr := mux.Vars(r)["id"]
	assetIDStr := mux.Vars(r)["asset_id"]
	portfolioID, err := strconv.Atoi(portfolioIDStr)
	if err != nil {
		http.Error(w, "Invalid portfolio ID", http.StatusBadRequest)
		return
	}
	assetID, err := strconv.Atoi(assetIDStr)
	if err != nil {
		http.Error(w, "Invalid asset ID", http.StatusBadRequest)
		return
	}
	asset, err := c.assetService.GetAssetByID(uint(assetID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if asset.PortfolioID != uint(portfolioID) {
		http.Error(w, "Asset does not belong to this portfolio", http.StatusForbidden)
		return
	}
	json.NewEncoder(w).Encode(asset)
}

func (c *AssetController) UpdateAsset(w http.ResponseWriter, r *http.Request) {
	assetIDStr := mux.Vars(r)["asset_id"]

	assetID, err := strconv.Atoi(assetIDStr)
	if err != nil {
		http.Error(w, "Invalid asset ID", http.StatusBadRequest)
		return
	}
	var asset models.Asset
	if err := json.NewDecoder(r.Body).Decode(&asset); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	asset.ID = uint(assetID)
	err = c.assetService.UpdateAsset(asset.ID, asset.Name, asset.Value, asset.CreatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *AssetController) DeleteAsset(w http.ResponseWriter, r *http.Request) {
	assetIDStr := mux.Vars(r)["asset_id"]
	assetID, err := strconv.Atoi(assetIDStr)
	if err != nil {
		http.Error(w, "Invalid asset ID", http.StatusBadRequest)
		return
	}
	err = c.assetService.DeleteAsset(uint(assetID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
