package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/marco-debortoli/mdeb-ledger/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Creating Transaction Categories
func (apiConfig *APIConfig) HandleCreateCategory(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	validTypes := []string{
		string(models.TYPE_CREDIT),
		string(models.TYPE_DEBIT),
		string(models.TYPE_TRANSFER),
	}

	if params.Type == "" || !slices.Contains(validTypes, params.Type) {
		respondWithError(w, http.StatusBadRequest, "Provided type must be one of 'credit', 'debit' or 'transfer'")
		return
	}

	cat, err := models.CreateTransactionCategory(apiConfig.DB, params.Name, models.CategoryType(params.Type))
	if err != nil {
		log.Println("failed to create transaction category:", err)
		respondWithError(w, http.StatusBadRequest, "failed to create transaction category")
		return
	}

	respondWithJSON(w, http.StatusCreated, cat)
}

// Retrieving a list of Transaction Categories
func (apiConfig *APIConfig) HandleListCategory(w http.ResponseWriter, r *http.Request) {
	activeParamString := r.URL.Query().Get("active")
	activeParamBool := false

	if activeParamString != "" {
		var err error
		activeParamBool, err = strconv.ParseBool(activeParamString)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "active URL parameter should be a boolean")
			return
		}
	}

	var cats []models.TransactionCategory
	if activeParamString == "" {
		cats = models.GetAllTransactionCategories(apiConfig.DB)
	} else if activeParamBool {
		cats = models.GetActiveTransactionCategories(apiConfig.DB)
	} else {
		cats = models.GetInactiveTransactionCategories(apiConfig.DB)
	}

	respondWithJSON(w, http.StatusOK, cats)
}

// Retrieving a single Transaction Category
func (apiConfig *APIConfig) HandleGetCategory(w http.ResponseWriter, r *http.Request) {
	catId, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid UUID format")
		return
	}

	cat, err := models.GetTransactionCategory(
		apiConfig.DB,
		catId,
	)
	if err != nil {
		respondWithError(
			w,
			http.StatusNotFound,
			fmt.Sprintf("Could not find category with ID: %v", catId),
		)
		return
	}

	respondWithJSON(w, http.StatusOK, cat)
}
