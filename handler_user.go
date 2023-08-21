package main

import (
	"encoding/json"
	"fmt"
	"github.com/AmelAbema/rss-aggregator/internal/database"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(1)
	type parameters struct {
		Name string `json:"name"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		log.Printf("Cannot create user: %v\n", err)
		return
	}
	respondWithJSON(w, http.StatusCreated, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUserByAPIKey(w http.ResponseWriter, r *http.Request) {
	fmt.Println(1)
	type parameters struct {
		Name string `json:"name"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		log.Printf("Cannot create user: %v\n", err)
		return
	}
	respondWithJSON(w, http.StatusCreated, databaseUserToUser(user))
}
