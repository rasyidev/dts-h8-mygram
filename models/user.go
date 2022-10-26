package models

import "time"

type User struct {
	ID        string    `json:"user" gorm:"primaryKey"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Age       uint      `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserCreatedResponseData struct {
	ID       string `json:"user" gorm:"primaryKey"`
	Age      uint   `json:"age"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (u *User) ToUserCreatedResponse() UserCreatedResponseData {
	return UserCreatedResponseData{
		ID:       u.ID,
		Age:      u.Age,
		Email:    u.Email,
		Username: u.Username,
	}
}
