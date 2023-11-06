package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	"github.com/marco-debortoli/mdeb-ledger/server/api"
)

func StartServer(apiConfig *api.APIConfig) {
	r := chi.NewRouter()

	log.Printf("Starting Server on Port %v\n", apiConfig.Port)

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Use(render.SetContentType(render.ContentTypeJSON))

	v1Router := chi.NewRouter()
	v1Router.Get("/health", apiConfig.HandlerHealth)

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
