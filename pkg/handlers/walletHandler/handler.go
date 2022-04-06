package walletHandler

import "gorm.io/gorm"

type walletHandler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) walletHandler {
	return walletHandler{db}
}
