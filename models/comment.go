package models

import "time"

type Comment struct {
	ID         string `json:"id" gorm:"primaryKey"`
	Message    string `json:"message"`
	UserID     string `json:"user_id"`
	User       User
	PhotoID    string `json:"photo_id"`
	Photo      Photo
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
