package models

import "gorm.io/gorm"

type ESingnature struct {
	gorm.Model
	Id        int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Signature string `json:"signature"`
	PublicKey string `json:"public_key"`
	UserId    int    `json:"user_id"`
}
