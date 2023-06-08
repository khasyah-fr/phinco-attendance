package routes

import (
	"github.com/gorilla/mux"
	"github.com/khasyah-fr/phinco-attendance/handlers"
)

func RegisterOnboardingRoutes(router *mux.Router) {
	router.HandleFunc("/onboarding", handlers.GetOnboardingHandler).Methods("GET")
}
