package routes

import (
	"github.com/gorilla/mux"
	"github.com/khasyah-fr/phinco-attendance/handlers"
)

func RegisterAuthenticationRoutes(router *mux.Router, ah *handlers.AuthenticationHandler) {
	router.HandleFunc("/register", ah.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", ah.LoginHandler).Methods("POST")
	router.HandleFunc("/forgot-password", ah.ForgotPasswordHandler).Methods("POST")
}
