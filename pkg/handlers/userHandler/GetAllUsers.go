package userHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type allUserInfo struct {
	Id      int
	Login   string
	Email   string
	Name    string
	Surname string
}

func (h userHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	var users []allUserInfo
	var user allUserInfo

	rows, e := h.DB.Table("user_accaunts").Joins("JOIN users on users.id = user_accaunts.user_id").
		Select("users.id, users.login, user_accaunts.email, user_accaunts.name, user_accaunts.surname").Rows()

	if e != nil {
		fmt.Println(w, "can't unmarshal: ", e.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
	defer rows.Close()

	i := 0
	for rows.Next() {
		// ScanRows scan a row into user
		h.DB.ScanRows(rows, &user)
		users = append(users, user)
		i++
	}

	rows.Close()

	json.NewEncoder(w).Encode(&users)
}
