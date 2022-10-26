package models

import "time"

type SocialMedia struct {
	ID             string `json:"id" gorm:"primaryKey"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserID         string `json:"user_id"`
	User           User
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
