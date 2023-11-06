package api

import (
	"encoding/json"
	"log"
	"net/http"

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
