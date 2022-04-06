package models

import "gorm.io/gorm"

type BalanceHistory struct {
	gorm.Model
	Id         int `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	EthBalance int `json:"eth_balance"`
	GasBalance int `json:"gas_balance"`
	WalletId   int `json:"wallet_Id"`
}
