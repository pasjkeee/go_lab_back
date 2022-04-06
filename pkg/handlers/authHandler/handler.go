package authHandler

import "gorm.io/gorm"

type authHandler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) authHandler {
	return authHandler{db}
}
