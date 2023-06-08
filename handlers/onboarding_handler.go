package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/khasyah-fr/phinco-attendance/entities"
	"github.com/khasyah-fr/phinco-attendance/helpers"
)

func GetOnboardingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var response entities.ResponseEntity

	var onboardingMessages = []entities.OnboardingMessage{
		{Id: 1, Title: "DIGITAL ABSENSI", Description: "Kehadiran sistem absensi digital merupakan penemuan yang mampu menggantikan pencatatan data kehadiran secara manual", URL: "https://example.com/images/absen.jpg"},
		{Id: 2, Title: "ATTENDANCE SYSTEM", Description: "Pengelolaan karyawan di era digital yang baik, menghasilkan karyawan terbaik pula, salah satunya absensi karyawan", URL: "https://example.com/images/orang.jpg"},
		{Id: 3, Title: "SELALU PAKAI MASKER", Description: "Guna mencegah penyebaran virus Covid-19, Pemerintah telah mengeluarkan kebijakan Physical Distancing serta kebijakan bekerja, belajar, dan beribadah dari rumah", URL: "https://example.com/images/masker.jpg"},
	}

	response.Status = http.StatusOK
	response.Data = onboardingMessages

	onboardingData, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helpers.DataToResponse(http.StatusInternalServerError, fmt.Sprintf("Failed to marshal response: %v", err)))
		return
	}

	w.Write(onboardingData)
}
