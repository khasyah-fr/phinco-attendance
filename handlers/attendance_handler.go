package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/khasyah-fr/phinco-attendance/entities"
	"github.com/khasyah-fr/phinco-attendance/helpers"
	"github.com/khasyah-fr/phinco-attendance/models"
	"github.com/khasyah-fr/phinco-attendance/repository"
)

type AttendanceHandler struct {
	attendanceRepo *repository.AttendanceRepository
	locationRepo   *repository.LocationRepository
}

func NewAttendanceHandler(attendanceRepo *repository.AttendanceRepository, locationRepo *repository.LocationRepository) *AttendanceHandler {
	return &AttendanceHandler{
		attendanceRepo: attendanceRepo,
		locationRepo:   locationRepo,
	}
}

var filters = map[string]int{"day": 1, "week": 1, "month": 1, "year": 1}

func (ah *AttendanceHandler) GetAllAttendancesByUserIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userId, _ := strconv.Atoi(mux.Vars(r)["user-id"])
	filter := r.URL.Query().Get("filter")

	if _, ok := filters[filter]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(helpers.DataToResponse(http.StatusBadRequest, "Wrong filter query"))
		return
	}

	attendances, err := ah.attendanceRepo.GetAllAttendancesByUserID(int64(userId), filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helpers.DataToResponse(http.StatusInternalServerError, "Failed to get all attendances by user id"))
		return
	}

	locations, err := ah.locationRepo.GetAllLocations()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helpers.DataToResponse(http.StatusInternalServerError, "Failed to get all locations"))
		return
	}

	w.Write(helpers.DataToResponse(http.StatusOK, combineAttendancesLocations(attendances, locations)))
}

func (ah *AttendanceHandler) CheckInHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var checkInRequest entities.CheckInRequest

	err := json.NewDecoder(r.Body).Decode(&checkInRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helpers.DataToResponse(http.StatusInternalServerError, "Failed to decode the request"))
		return
	}

	err = ah.attendanceRepo.CheckIn(checkInRequest.UserId, checkInRequest.LocationId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helpers.DataToResponse(http.StatusInternalServerError, "Failed to check-in"))
		return
	}

	w.Write(helpers.DataToResponse(http.StatusOK, "Checked in successfully"))
}

func (ah *AttendanceHandler) CheckOutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var checkOutRequest entities.CheckOutRequest

	err := json.NewDecoder(r.Body).Decode(&checkOutRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helpers.DataToResponse(http.StatusInternalServerError, "Failed to decode the request"))
		return
	}

	err = ah.attendanceRepo.CheckOut(checkOutRequest.UserId, checkOutRequest.LocationId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helpers.DataToResponse(http.StatusInternalServerError, "Failed to check-out"))
		return
	}

	w.Write(helpers.DataToResponse(http.StatusOK, "Checked out successfully"))
}

func combineAttendancesLocations(attendances []models.Attendance, locations []models.Location) []entities.AttendanceResponse {
	var ar []entities.AttendanceResponse

	locationMap := make(map[int64]models.Location)

	for _, location := range locations {
		locationMap[location.Id] = location
	}

	for _, attendance := range attendances {
		location, exist := locationMap[attendance.LocationId]
		if !exist {
			continue
		}

		ar = append(ar, entities.AttendanceResponse{Type: attendance.Type, Name: location.Name, Address: location.Address, Url: location.Url, AttendTime: attendance.AttendTime})
	}

	return ar
}
