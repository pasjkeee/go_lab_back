package models

import (
	"gorm.io/gorm"
	"time"
)

type Transactions struct {
	gorm.Model
	Id                int             `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	DateTime          time.Time       `json:"date_time" gorm:"type:time"`
	EntityId          int             `json:"entity_id"`
	TransactionMetaId TransactionMeta `gorm:"ForeignKey:TransactionId;constraint:OnUpdate:CASCADE,OnDelete:Cascade;"`
}
