package handlers

import (
	"net/http"
	"simple_api/utils"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Health check endpoint called")

	utils.JsonResponse(w, http.StatusOK, "Service is healthy", nil)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Getting users")
	// Example data
	users := []map[string]interface{}{
		{
			"id":    1,
			"name":  "John Doe",
			"email": "john@example.com",
		},
		{
			"id":    2,
			"name":  "Jane Doe",
			"email": "jane@example.com",
		},
	}
	utils.JsonResponse(w, http.StatusOK, "Users retrieved successfully", users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Creating user")
	// Example create user logic
	newUser := map[string]interface{}{
		"id":    3,
		"name":  "New User",
		"email": "new@example.com",
	}
	utils.JsonResponse(w, http.StatusCreated, "User created successfully", newUser)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	log.Info().Str("id", id).Msg("Getting user")

	// Example single user data
	user := map[string]interface{}{
		"id":    id,
		"name":  "John Doe",
		"email": "john@example.com",
	}
	utils.JsonResponse(w, http.StatusOK, "User retrieved successfully", user)
}
