package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/marco-debortoli/mdeb-ledger/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Creating Accounts
func (apiConfig *APIConfig) HandleCreateAccount(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name            string `json:"name"`
		Type            string `json:"type"`
		IncludeNetWorth bool   `json:"include_net_worth"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	validTypes := []string{
		string(models.TYPE_ACCOUNT_DEBIT),
		string(models.TYPE_ACCOUNT_CREDIT),
		string(models.TYPE_ACCOUNT_INVESTMENT),
	}

	if params.Type == "" || !slices.Contains(validTypes, params.Type) {
		respondWithError(
			w,
			http.StatusBadRequest,
			"Provided type must be one of 'credit', 'debit' or 'investment'",
		)
		return
	}

	acct, err := models.CreateAccount(
		apiConfig.DB,
		params.Name,
		models.AccountType(params.Type),
		params.IncludeNetWorth,
		primitive.NewDateTimeFromTime(time.Now()),
	)
	if err != nil {
		log.Println("failed to create account:", err)
		respondWithError(w, http.StatusBadRequest, "failed to create account")
		return
	}

	respondWithJSON(w, http.StatusCreated, acct)
}

// Account actions
func (apiConfig *APIConfig) HandleSetAccountCurrentValue(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL
	accountId, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid UUID format")
		return
	}

	// Now get the Account from the provided ID (or return a 404)
	account, err := models.GetAccount(apiConfig.DB, accountId)

	if err != nil {
		respondWithError(
			w,
			http.StatusNotFound,
			fmt.Sprintf("Could not find account with ID: %v", accountId),
		)

		return
	}

	// Decode the parameters
	type parameters struct {
		Value float64 `json:"value"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	// Now update the account value
	err = models.SetAccountCurrentValue(apiConfig.DB, &account, params.Value, time.Now())

	if err != nil {
		respondWithError(
			w,
			http.StatusBadRequest,
			fmt.Sprintf("Failed to set account value: %v", err),
		)
		return
	}

	respondWithJSON(w, http.StatusOK, account)
}
