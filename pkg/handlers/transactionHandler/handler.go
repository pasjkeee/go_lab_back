package transactionHandler

import "gorm.io/gorm"

type transactionHandler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) transactionHandler {
	return transactionHandler{db}
}
