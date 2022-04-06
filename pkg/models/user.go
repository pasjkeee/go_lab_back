package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id            *int                     `json:"id" gorm:"primaryKey;autoIncrement"`
	Login         string                   `json:"login"`
	Password      string                   `json:"password"`
	ESingnatureId *ESingnature             `gorm:"ForeignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:Cascade"`
	UserAccaunt   *UserAccaunt             `gorm:"ForeignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	OwnerIds      *Nft                     `gorm:"ForeignKey:OwnerId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatorIds    *Nft                     `gorm:"ForeignKey:CreatorId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	AuthStatuses  *[]AuthStatus            `gorm:"many2many:user_auth_status"`
	Wallets       *[]Wallet                `gorm:"many2many:user_wallets"`
	Permissions   *[]UserPermissionsGroups `gorm:"many2many:user_permissions;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (user User) UserWalletTableName() string {
	return "user_wallets"
}

func (user User) TableName() string {
	return "users"
}

func (user *User) FindUserForWalletId(id int, h DBInterfase) {
	h.First(user, User{Wallets: &[]Wallet{{Id: id}}})
}

func (user *User) FindUserByLogin(login string, h DBInterfase) {
	h.First(user, User{Login: login})
}
