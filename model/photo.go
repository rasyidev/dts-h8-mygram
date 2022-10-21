package model

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	Title    string
	Caption  string
	PhotoURL string
	UserID   string
}
