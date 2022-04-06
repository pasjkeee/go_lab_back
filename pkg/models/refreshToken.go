package models

import "gorm.io/gorm"

type RefreshToken struct {
	gorm.Model
	Token string `json:"token" gorm:"primaryKey;AUTO_INCREMENT"`
	Login string `json:"login"`
}

func (rt *RefreshToken) Find(token string, h DBInterfase) {
	h.First(rt, RefreshToken{Token: token})
}

func (rt *RefreshToken) Add(h DBInterfase) {
	h.Create(rt)
}

func (rt *RefreshToken) Delete(h DBInterfase) {
	h.Delete(rt, RefreshToken{Token: rt.Token})
}

func (rt *RefreshToken) DeleteByLogin(login string, h DBInterfase) {
	h.Delete(rt, RefreshToken{Login: login})
}
