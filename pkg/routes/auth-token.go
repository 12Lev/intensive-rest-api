package routes

import (
	"encoding/json"
	"intensive-rest-api/pkg/auth"
	"net/http"
)

// PostAuth godoc
// @Summary Авторизация пользователя.
// @Description Авторизация пользователя.
// @ID routes-auth-user-profile
// @Tags Авторизация
// @Accept json
// @Param body body auth.AuthRequest true  "Запрос"
// @Produce json
// @Success 200 {object} auth.AuthResponse
// @Failure 400 "Error: Bad Request"
// @Failure 500
// @Router /api/v1/token [post]
func PostAuth(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		(w).WriteHeader(http.StatusOK)
		return
	}
	var inputRequest auth.AuthRequest
	err := json.NewDecoder(r.Body).Decode(&inputRequest)
	if err != nil {
		(w).WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{
			"error": err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	response, err := inputRequest.Auth()
	if err != nil {
		(w).WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{
			"error": err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	(w).WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
