package models

import "gorm.io/gorm"

type UserAccaunt struct {
	gorm.Model
	Id      int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	UserId  int    `json:"user_id"`
}
