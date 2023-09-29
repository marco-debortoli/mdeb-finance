package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func startServer() {
	r := chi.NewRouter()

	fmt.Printf("Starting Server on Port 3000")

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Use(render.SetContentType(render.ContentTypeJSON))

	// r.Route()

	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe(":3000", r)
}
