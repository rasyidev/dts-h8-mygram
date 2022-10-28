package models

import "time"

type Comment struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Message   string `json:"message" validate:"required"`
	UserID    string `json:"user_id"`
	User      User
	PhotoID   string `json:"photo_id"`
	Photo     Photo
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostCommentRequest struct {
	Message string `json:"message" validate:"required"`
	PhotoID string `json:"photo_id" validate:"required"`
}

type UserCommentResponse struct {
	ID       string `json:"user"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PhotoCommentResponse struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" validate:"required,url"`
	UserID   string `json:"user_id"`
}

type PostCommentResponse struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	UserID    string `json:"user_id"`
	User      UserCommentResponse
	PhotoID   string `json:"photo_id"`
	Photo     PhotoCommentResponse
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Comment) ToPostCommentResponse() PostCommentResponse {
	return PostCommentResponse{
		ID:        c.ID,
		Message:   c.Message,
		PhotoID:   c.PhotoID,
		UserID:    c.UserID,
		CreatedAt: c.CreatedAt,
	}
}
