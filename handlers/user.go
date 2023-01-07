package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cecardev/go-rest-server/models"
	"github.com/cecardev/go-rest-server/repository"
	"github.com/cecardev/go-rest-server/server"
	"github.com/segmentio/ksuid"
)

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
}

func SignUpHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = SignUpRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		idInt, _ := strconv.ParseInt(id.String(), 10, 0)
		var user = models.User{
			Email:    request.Email,
			Password: request.Password,
			Id:       idInt,
		}

		err = repository.InsertUser(r.Context(), &user)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SignUpResponse{
			Id:    user.Id,
			Email: user.Email,
		})

	}
}
