package models

import "time"

type Location struct {
	Id        int64     `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Address   string    `gorm:"not null" json:"address"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
