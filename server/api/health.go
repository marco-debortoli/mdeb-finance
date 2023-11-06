package api

import "net/http"

func (apiConfig *APIConfig) HandleHealth(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, struct{}{})
}
