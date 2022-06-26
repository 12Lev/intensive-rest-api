package routes

import (
	"intensive-rest-api/pkg/auth"
	"intensive-rest-api/pkg/utils"
	"net/http"
	"strings"
)

func setupResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, token, refresh-token")
}

func AuthRequired(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		token := r.Header.Get("token")
		if len(strings.TrimSpace(token)) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if !auth.Validate(token) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(utils.ToJsonBytes(map[string]string{"state": "Token is not validate"}))
		}
	}
}

// GetOnlyAuthorized godoc
// @Summary Получение данных, только для авторизованных пользователей.
// @Description Получение данных.
// @ID routes-only-authorized-user
// @Security ApiKeyAuth
// @Tags Данные
// @Accept json
// @Produce json
// @Success 200 "OK"
// @Failure 401 "Error: Unauthorized"
// @Failure 500
// @Router /api/v1/get-data [get]
func GetOnlyAuthorized(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		(w).WriteHeader(http.StatusOK)
		return
	}
	(w).WriteHeader(http.StatusOK)
}
