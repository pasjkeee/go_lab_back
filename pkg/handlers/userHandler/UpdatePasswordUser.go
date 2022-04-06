package userHandler

import (
	"awesomeProject/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func (h userHandler) UpdatePasswordUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "err %q %q\n", err, err.Error())
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var requsetUser models.User

	err = json.Unmarshal(body, &requsetUser)
	if err != nil {
		fmt.Println(w, "can't unmarshal: ", err.Error())
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.DB.First(&user, models.User{Login: params["login"]})
	user.Password = requsetUser.Password
	h.DB.Save(&user)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}
