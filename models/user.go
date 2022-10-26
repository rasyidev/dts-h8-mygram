package models

import "time"

type User struct {
	ID         string    `json:"user" gorm:"primaryKey"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Age        uint      `json:"age"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
