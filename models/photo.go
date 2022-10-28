package models

import "time"

type Photo struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Title     string `json:"title" validate:"required"`
	Caption   string `json:"caption"`
	PhotoUrl  string `json:"photo_url" validate:"required,url"`
	UserID    string `json:"user_id"`
	User      User
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PhotoPostRequest struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" validate:"required,url"`
}

func (p *Photo) ToPhotoPostRequest() PhotoPostRequest {
	return PhotoPostRequest{
		Title:    p.Title,
		Caption:  p.Caption,
		PhotoUrl: p.PhotoUrl,
	}
}

type PhotoPostedResponse struct {
	Title     string    `json:"title" validate:"required"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url" validate:"required,url"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Photo) ToPhotoPostedResponse() PhotoPostedResponse {
	return PhotoPostedResponse{
		Title:     p.Title,
		Caption:   p.Caption,
		PhotoUrl:  p.PhotoUrl,
		UserID:    p.UserID,
		CreatedAt: p.CreatedAt,
	}
}

type UserSecured struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type GetAllPhotosReponse struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" validate:"required"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url" validate:"required,url"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      UserSecured
}

func (p *Photo) ToGetAllPhotosResponse() GetAllPhotosReponse {
	return GetAllPhotosReponse{
		ID:       p.ID,
		Title:    p.Title,
		Caption:  p.Caption,
		PhotoUrl: p.PhotoUrl,
		UserID:   p.UserID,
		User: UserSecured{
			Email:    p.User.Email,
			Username: p.User.Username,
		},
	}
}

type UpdatedPhotosReponse struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" validate:"required"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url" validate:"required,url"`
	UserID    string    `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Photo) ToUpdatedPhotosResponse() UpdatedPhotosReponse {
	return UpdatedPhotosReponse{
		ID:        p.ID,
		Title:     p.Title,
		Caption:   p.Caption,
		PhotoUrl:  p.PhotoUrl,
		UserID:    p.UserID,
		UpdatedAt: p.UpdatedAt,
	}
}
