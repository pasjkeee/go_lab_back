package userHandler

import (
	"awesomeProject/pkg/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (h userHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "err %q %q\n", err, err.Error())
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println(w, "can't unmarshal: ", err.Error())
	}

	resUser := models.User{Login: user.Login, Password: user.Password}
	res := h.DB.Create(&resUser)

	if res.Error != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&res.Error)
		fmt.Println("error")
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&resUser)
}
