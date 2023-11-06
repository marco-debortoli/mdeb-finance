package api

import "net/http"

func (apiConfig *APIConfig) HandlerHealth(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, struct{}{})
}
