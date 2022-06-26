package routes

import (
	"encoding/json"
	"intensive-rest-api/pkg/auth"
	"net/http"
)

// PostAddUser godoc
// @Summary Добавление пользователя.
// @Description Добавление пользователя.
// @ID routes-add-user
// @Tags Пользователи
// @Accept json
// @Param body body auth.CreateUserRequest true  "Запрос"
// @Produce json
// @Success 200 {object} auth.CreateUserResponse
// @Failure 400 "Error: Bad Request"
// @Failure 500
// @Router /api/v1/add-user [post]
func PostAddUser(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		//здесь можно сделать проверку
		(w).WriteHeader(http.StatusOK)
		return
	}
	var inputRequest auth.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&inputRequest)
	if err != nil {
		(w).WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{
			"error": err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	response, err := inputRequest.AddOrUpdateUser()
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
