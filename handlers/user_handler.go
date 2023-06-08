package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/khasyah-fr/phinco-attendance/helpers"
	"github.com/khasyah-fr/phinco-attendance/repository"
)

type UserHandler struct {
	userRepo *repository.UserRepository
}

func NewUserHandler(userRepo *repository.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

func (ah *UserHandler) GetOneUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	existingUser, err := ah.userRepo.GetUserByUsername(mux.Vars(r)["username"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helpers.DataToResponse(http.StatusInternalServerError, "Failed to fetch the user"))
		return
	}

	if existingUser == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(helpers.DataToResponse(http.StatusBadRequest, "Username is incorrect"))
		return
	}

	w.Write(helpers.DataToResponse(http.StatusOK, existingUser))
}
