package models

import "gorm.io/gorm"

type DBInterfase interface {
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Create(value interface{}) (tx *gorm.DB)
	Delete(value interface{}, conds ...interface{}) (tx *gorm.DB)
	Where(value interface{}, conds ...interface{}) (tx *gorm.DB)
}
type Wallet struct {
	gorm.Model
	Id                int              `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	EthBalance        float64          `json:"eth_balance"`
	GasBalance        float64          `json:"gas_balance"`
	WalletTrasactions []Transactions   `gorm:"many2many:wallet_transactions"`
	BalanceHistoryIds []BalanceHistory `gorm:"ForeignKey:WalletId;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
}

func (w *Wallet) FindWallet(id int, h DBInterfase) {
	h.First(w, Wallet{Id: id})
}

func (w *Wallet) ChangeEtheriumBalance(value float64) {
	w.EthBalance = w.EthBalance + value
}
