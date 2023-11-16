package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/marco-debortoli/mdeb-ledger/server/api"
)

func StartServer(apiConfig *api.APIConfig) {
	r := chi.NewRouter()

	log.Printf("Starting Server on Port %v\n", apiConfig.Port)

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	v1Router := chi.NewRouter()
	// Server Maintenance
	v1Router.Get("/health", apiConfig.HandleHealth)

	// Categories
	v1Router.Post("/categories", apiConfig.HandleCreateCategory)
	v1Router.Get("/categories", apiConfig.HandleListCategory)
	v1Router.Get("/categories/{id}", apiConfig.HandleGetCategory)

	// Transactions
	v1Router.Post("/transactions", apiConfig.HandleCreateTransaction)
	v1Router.Get("/transactions", apiConfig.HandleListTransaction)
	v1Router.Delete("/transactions/{id}", apiConfig.HandleDeleteTransaction)
	v1Router.Put("/transactions/{id}", apiConfig.HandleUpdateTransaction)

	// Accounts
	v1Router.Post("/accounts", apiConfig.HandleCreateAccount)
	v1Router.Post("/accounts/{id}/set_value", apiConfig.HandleSetAccountValue)

	r.Mount("/api/v1", v1Router)

	srv := &http.Server{
		Handler: r,
		Addr:    ":" + apiConfig.Port,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("Server stopped:", err)
	}
}
