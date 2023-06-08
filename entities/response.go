package entities

import "time"

type ResponseEntity struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type AttendanceResponse struct {
	Type       string    `json:"type"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	Url        string    `json:"url"`
	AttendTime time.Time `json:"attend_time"`
}
