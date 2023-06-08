package models

import "time"

type Attendance struct {
	Id         int64     `gorm:"primary_key" json:"id"`
	UserId     int64     `gorm:"not null" json:"user_id"`
	LocationId int64     `gorm:"not null" json:"location_id"`
	Type       string    `gorm:"not null" json:"type"`
	AttendTime time.Time `json:"attend_time"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
