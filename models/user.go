package models

import "time"

type User struct {
	ID        string    `json:"user" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique" validate:"required"`
	Email     string    `json:"email" gorm:"unique" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=6"`
	Age       uint      `json:"age" validate:"required,min=9"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserCreatedResponseData struct {
	ID       string `json:"user"`
	Age      uint   `json:"age"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserUpdatedResponseData struct {
	ID        string    `json:"user"`
	Age       uint      `json:"age"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) ToUserCreatedResponse() UserCreatedResponseData {
	return UserCreatedResponseData{
		ID:       u.ID,
		Age:      u.Age,
		Email:    u.Email,
		Username: u.Username,
	}
}

func (u *User) ToUserUpdatedResponse() UserUpdatedResponseData {
	return UserUpdatedResponseData{
		ID:        u.ID,
		Age:       u.Age,
		Email:     u.Email,
		Username:  u.Username,
		UpdatedAt: u.UpdatedAt,
	}
}
