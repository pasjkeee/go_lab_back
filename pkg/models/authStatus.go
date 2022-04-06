package models

import "gorm.io/gorm"

type AuthStatus struct {
	gorm.Model
	Id           int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Token        string `json:"token"`
	AuthLocation string `json:"auth_location"`
}
