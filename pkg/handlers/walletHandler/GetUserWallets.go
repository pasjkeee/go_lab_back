package walletHandler

import (
	"awesomeProject/pkg/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func (h walletHandler) GetUserWallets(w http.ResponseWriter, r *http.Request) {
	var wallets []models.Wallet

	params := mux.Vars(r)

	h.DB.Joins("JOIN user_wallets on user_wallets.wallet_id=wallets.id").
		Joins("JOIN users on users.id=user_wallets.user_id").
		Where("users.id in (?) ", []string{params["id"]}).Find(&wallets)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&wallets)

}
