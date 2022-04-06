package walletHandler

import (
	"awesomeProject/pkg/models"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (h walletHandler) AddWallet(w http.ResponseWriter, r *http.Request) {
	var wallet models.Wallet
	var user models.User

	params := mux.Vars(r)
	h.DB.First(&user, models.User{Login: params["login"]})

	user.Wallets = &[]models.Wallet{wallet}

	h.DB.Save(&user)

	fmt.Println(user.Wallets)

}
