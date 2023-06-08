package routes

import (
	"github.com/gorilla/mux"
	"github.com/khasyah-fr/phinco-attendance/handlers"
	"github.com/khasyah-fr/phinco-attendance/middlewares"
)

func RegisterAttendanceRoutes(router *mux.Router, ah *handlers.AttendanceHandler) {
	router.HandleFunc("/check-in", middlewares.JwtAuthentication(ah.CheckInHandler)).Methods("POST")
	router.HandleFunc("/check-out", middlewares.JwtAuthentication(ah.CheckOutHandler)).Methods("POST")
	router.HandleFunc("/attendances/{user-id}", middlewares.JwtAuthentication(ah.GetAllAttendancesByUserIdHandler)).Methods("GET")
}
