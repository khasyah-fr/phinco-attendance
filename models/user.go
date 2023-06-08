package models

import "time"

type User struct {
	Id        int64     `gorm:"primary_key" json:"id"`
	Username  string    `gorm:"not null;unique" json:"username"`
	Fullname  string    `gorm:"not null" json:"fullname"`
	Password  string    `gorm:"not null" json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
