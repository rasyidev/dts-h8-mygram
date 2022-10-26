package models

import "time"

type Comment struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Message   string `json:"message"`
	UserID    string `json:"user_id"`
	User      User
	PhotoID   string `json:"photo_id"`
	Photo     Photo
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
