package routes

import (
	"github.com/gorilla/mux"
	"github.com/khasyah-fr/phinco-attendance/handlers"
	"github.com/khasyah-fr/phinco-attendance/middlewares"
)

func RegisterLocationRoutes(router *mux.Router, lh *handlers.LocationHandler) {
	router.HandleFunc("/locations", middlewares.JwtAuthentication(lh.GetAllLocationsHandler)).Methods("GET")
}
