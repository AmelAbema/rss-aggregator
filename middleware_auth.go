package main

import (
	"fmt"
	"github.com/AmelAbema/rss-aggregator/internal/auth"
	"github.com/AmelAbema/rss-aggregator/internal/database"
	"net/http"
)

type authedHeader func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHeader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Cannot get API key: %v", err))
			return
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Cannot get User with API key: %v", err))
			return
		}

		handler(w, r, user)
	}
}
