package models

import "time"

type Photo struct {
	ID         string `json:"id" gorm:"primaryKey"`
	Title      string `json:"title"`
	Caption    string `json:"caption"`
	PhotoUrl   string `json:"photo_url"`
	UserID     string `json:"user_id"`
	User       User
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
