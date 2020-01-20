package handlers

import "net/http"

func HandleCors(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Access-Control-Request-Method") != "" {
		// Set CORS headers
		header := w.Header()
		header.Set("Access-Control-Allow-Methods", r.Header.Get("Allow"))
		header.Set("Access-Control-Allow-Origin", "*")
	}

	// Adjust status code to 204
	w.WriteHeader(http.StatusNoContent)
}
