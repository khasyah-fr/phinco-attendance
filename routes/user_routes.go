package routes

import (
	"github.com/gorilla/mux"
	"github.com/khasyah-fr/phinco-attendance/handlers"
	"github.com/khasyah-fr/phinco-attendance/middlewares"
)

func RegisterUserRoutes(router *mux.Router, uh *handlers.UserHandler) {
	router.HandleFunc("/users/{username}", middlewares.JwtAuthentication(uh.GetOneUserHandler)).Methods("GET")
}
