package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/khasyah-fr/phinco-attendance/config"
	"github.com/khasyah-fr/phinco-attendance/database"
	"github.com/khasyah-fr/phinco-attendance/handlers"
	"github.com/khasyah-fr/phinco-attendance/repository"
	"github.com/khasyah-fr/phinco-attendance/routes"
)

func main() {
	// config
	config := config.GetConfig()

	db, err := database.InitializeDB(config.DB)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	attendanceRepo := repository.NewAttendanceRepository(db)
	locationRepo := repository.NewLocationRepository(db)
	userRepo := repository.NewUserRepository(db)

	attendanceHandler := handlers.NewAttendanceHandler(attendanceRepo, locationRepo)
	authenticationHandler := handlers.NewAuthenticationHandler(userRepo)
	locationHandler := handlers.NewLocationHandler(locationRepo)
	userHandler := handlers.NewUserHandler(userRepo)

	router := mux.NewRouter()

	routes.RegisterAttendanceRoutes(router, attendanceHandler)
	routes.RegisterAuthenticationRoutes(router, authenticationHandler)
	routes.RegisterLocationRoutes(router, locationHandler)
	routes.RegisterOnboardingRoutes(router)
	routes.RegisterUserRoutes(router, userHandler)

	log.Println("Server listening on port 8000")
	err = http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}
