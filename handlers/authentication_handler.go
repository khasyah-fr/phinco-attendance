package handlers

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/khasyah-fr/phinco-attendance/entities"
	"github.com/khasyah-fr/phinco-attendance/helpers"
	"github.com/khasyah-fr/phinco-attendance/models"
	"github.com/khasyah-fr/phinco-attendance/repository"
)

type AuthenticationHandler struct {
	userRepo *repository.UserRepository
}

func NewAuthenticationHandler(userRepo *repository.UserRepository) *AuthenticationHandler {
	return &AuthenticationHandler{
		userRepo: userRepo,
	}
}

func (ah *AuthenticationHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var registerRequest entities.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&registerRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helpers.DataToResponse(http.StatusInternalServerError, "Failed to decode the request"))
		return
	}

	if registerRequest.Password != registerRequest.RepeatPassword {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(helpers.DataToResponse(http.StatusBadRequest, "Password is mismatched"))
		return
	}

	existingUser, err := ah.userRepo.GetUserByUsername(registerRequest.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helpers.DataToResponse(http.StatusInternalServerError, "Failed to fetch the user"))
		return
	}

	if existingUser != nil {
		w.WriteHeader(http.StatusConflict)
		w.Write(helpers.DataToResponse(http.StatusConflict, "Username is already taken"))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helpers.DataToResponse(http.StatusInternalServerError, "Failed to hash password"))
		return
	}
	registerRequest.Password = string(hashedPassword)

	err = ah.userRepo.CreateUser(registerRequestToUser(&registerRequest))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helpers.DataToResponse(http.StatusInternalServerError, "Failed to create the user"))
		return
	}

	w.Write(helpers.DataToResponse(http.StatusOK, "User registered successfully"))
}

func (ah *AuthenticationHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var loginRequest entities.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helpers.DataToResponse(http.StatusInternalServerError, "Failed to decode the request"))
		return
	}

	existingUser, err := ah.userRepo.GetUserByUsername(loginRequest.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helpers.DataToResponse(http.StatusInternalServerError, "Failed to fetch the user"))
		return
	}

	if existingUser == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(helpers.DataToResponse(http.StatusBadRequest, "Username or password is incorrect"))
		return
	}

	if !comparePasswords(existingUser.Password, loginRequest.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(helpers.DataToResponse(http.StatusUnauthorized, "Username or password is incorrect"))
		return
	}

	token, err := helpers.GenerateJWTToken(loginRequest.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helpers.DataToResponse(http.StatusInternalServerError, "Failed to generate jwt token"))
		return
	}

	w.Header().Set("Authorization", "Bearer "+token)

	w.Write(helpers.DataToResponse(http.StatusOK, "Login successful"))
}

func (ah *AuthenticationHandler) ForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var forgotPasswordRequest entities.ForgotPasswordRequest
	err := json.NewDecoder(r.Body).Decode(&forgotPasswordRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helpers.DataToResponse(http.StatusInternalServerError, "Failed to decode the request"))
		return
	}

	if forgotPasswordRequest.Password != forgotPasswordRequest.RepeatPassword {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(helpers.DataToResponse(http.StatusBadRequest, "Password is mismatched"))
		return
	}

	existingUser, err := ah.userRepo.GetUserByUsername(forgotPasswordRequest.Username)
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(forgotPasswordRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helpers.DataToResponse(http.StatusInternalServerError, "Failed to hash password"))
		return
	}
	forgotPasswordRequest.Password = string(hashedPassword)

	err = ah.userRepo.UpdatePasswordByUsername(forgotPasswordRequest.Username, forgotPasswordRequest.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helpers.DataToResponse(http.StatusInternalServerError, "Failed to update the password"))
		return
	}

	w.Write(helpers.DataToResponse(http.StatusOK, "Password changed successfully"))
}

func registerRequestToUser(rr *entities.RegisterRequest) *models.User {
	return &models.User{
		Username: rr.Username,
		Fullname: rr.Fullname,
		Password: rr.Password,
	}
}

func comparePasswords(expected string, input string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(expected), []byte(input))
	if err != nil {
		// Password does not match
		return false
	}

	// Password matches
	return true
}
