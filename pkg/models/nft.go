package models

import "gorm.io/gorm"

type Nft struct {
	gorm.Model
	Id           int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	MetaId       int    `json:"meta_id"`
	LocationLink string `json:"location_link"`
	OwnerId      int    `json:"owner_id"`
	CreatorId    int    `json:"creator_id"`
}
