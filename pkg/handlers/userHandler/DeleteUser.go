package userHandler

import (
	"awesomeProject/pkg/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func (h userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	h.DB.First(&user, models.User{Login: params["login"]})
	h.DB.Delete(&user)
	json.NewEncoder(w).Encode(&user)
}
