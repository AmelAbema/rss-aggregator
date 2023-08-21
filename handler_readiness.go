package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, _ *http.Request) {
	type response struct {
		Status string `json:"status"`
	}
	respondWithJSON(w, 200, response{Status: "ok"})
}
