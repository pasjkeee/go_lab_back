package userHandler

import (
	"gorm.io/gorm"
)

type userHandler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) userHandler {
	return userHandler{db}
}
