package walletHandler

import (
	"awesomeProject/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (h walletHandler) GetWalletTransactions(w http.ResponseWriter, r *http.Request) {
	var transactions []models.Transactions
	params := mux.Vars(r)

	h.DB.Joins("JOIN wallet_transactions on wallet_transactions.transactions_id=transactions.id").
		Joins("JOIN wallets on wallets.id=wallet_transactions.wallet_id").
		Where("wallets.id in (?) ", []string{params["id"]}).Find(&transactions)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&transactions)

	fmt.Println(transactions)
}
