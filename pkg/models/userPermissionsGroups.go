package models

import "gorm.io/gorm"

type UserPermissionsGroups struct {
	gorm.Model
	Id    int    `json:"id" gorm:"primaryKey"`
	Value string `json:"value"`
	//Users []User `gorm:"ForeignKey:Permissions;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
