package handlers

import (
	"net/http"

	"github.com/khasyah-fr/phinco-attendance/helpers"
	"github.com/khasyah-fr/phinco-attendance/repository"
)

type LocationHandler struct {
	locationRepo *repository.LocationRepository
}

func NewLocationHandler(locationRepo *repository.LocationRepository) *LocationHandler {
	return &LocationHandler{
		locationRepo: locationRepo,
	}
}

func (lh *LocationHandler) GetAllLocationsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	allLocations, err := lh.locationRepo.GetAllLocations()
	if allLocations == nil || err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helpers.DataToResponse(http.StatusInternalServerError, "Failed to fetch all locations"))
		return
	}

	w.Write(helpers.DataToResponse(http.StatusOK, allLocations))
}
