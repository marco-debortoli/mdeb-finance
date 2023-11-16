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
func (apiConfig *APIConfig) HandleSetAccountValue(w http.ResponseWriter, r *http.Request) {
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
		Value float64             `json:"value"`
		Date  *primitive.DateTime `json:"date"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	// TODO - Here we should add validation to Date to ensure that it is the beginning of a month
	// This would limit it to one account value per month
	// For now we will just assume that the frontend is providing the proper date

	// Date must be provided
	if params.Date == nil {
		respondWithError(
			w,
			http.StatusBadRequest,
			"Date must be provided",
		)
		return
	}

	// Now update the account value
	err = models.SetAccountValue(apiConfig.DB, &account, params.Value, *params.Date)

	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			fmt.Sprintf("Failed to set account value: %v", err),
		)
		return
	}

	respondWithJSON(w, http.StatusOK, account)
}
