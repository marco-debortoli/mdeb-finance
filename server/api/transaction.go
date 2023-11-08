package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/marco-debortoli/mdeb-ledger/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Creating Transactions
func (apiConfig *APIConfig) HandleCreateTransaction(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name     string             `json:"name"`
		Amount   float64            `json:"amount"`
		Date     primitive.DateTime `json:"date"`
		Category primitive.ObjectID `json:"category"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	// Need to validate that all fields are provided and that the provided values are correct
	if params.Date.Time().Unix() == 0 {
		respondWithError(w, http.StatusBadRequest, "date must be provided")
		return
	}

	if params.Name == "" {
		respondWithError(w, http.StatusBadRequest, "name must be provided")
		return
	}

	if params.Amount <= 0.0 {
		respondWithError(w, http.StatusBadRequest, "amount must be provided and greater than zero")
		return
	}

	cat, err := models.GetTransactionCategory(apiConfig.DB, params.Category)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "provided category does not exist")
		return
	}

	transaction, err := models.CreateTransaction(apiConfig.DB, params.Name, params.Date, params.Amount, cat)
	if err != nil {
		log.Println("Failed to create transaction:", err)
		respondWithError(w, http.StatusBadRequest, "failed to create transaction")
		return
	}

	respondWithJSON(w, http.StatusCreated, transaction)
}

// List Transactions
func (apiConfig *APIConfig) HandleListTransaction(w http.ResponseWriter, r *http.Request) {
	startDateParam := r.URL.Query().Get("start_date")
	endDateParam := r.URL.Query().Get("end_date")

	var err error
	var startDateFilter time.Time
	var endDateFilter time.Time

	if startDateParam != "" {
		startDateFilter, err = time.Parse(time.RFC3339, startDateParam)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "start_date URL parameter should be a valid ISO-8601 datetime string")
			return
		}
	}

	if endDateParam != "" {
		endDateFilter, err = time.Parse(time.RFC3339, endDateParam)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "end_date URL parameter should be a valid ISO-8601 datetime string")
			return
		}
	}

	var transactionFilter *models.TransactionFilter

	if !startDateFilter.IsZero() || !endDateFilter.IsZero() {
		transactionFilter = &models.TransactionFilter{
			StartDate: startDateFilter,
			EndDate:   endDateFilter,
		}
	}

	transactions := models.GetTransactionList(apiConfig.DB, transactionFilter)
	respondWithJSON(w, http.StatusOK, transactions)
}

// Delete Transaction
func (apiConfig *APIConfig) HandleDeleteTransaction(w http.ResponseWriter, r *http.Request) {
	trId, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid UUID format")
		return
	}

	err = models.DeleteTransaction(apiConfig.DB, trId)
	if err != nil {
		respondWithError(
			w,
			http.StatusBadRequest,
			fmt.Sprintf("Could not delete transaction: %v", err),
		)
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}
