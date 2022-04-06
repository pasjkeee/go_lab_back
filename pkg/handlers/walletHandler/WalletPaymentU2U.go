package walletHandler

import (
	"awesomeProject/pkg/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type walletPaymentU2URequestBody struct {
	UserId           int     `json:"user_id"`
	SenderWalletId   int     `json:"sender_wallet_id"`
	ReceiverWalletId int     `json:"receiver_wallet_id"`
	PaymentValue     float64 `json:"payment_value"`
	Taxes            int     `json:"taxes"`
}

func fillTransaction(transaction *models.Transactions, reqBody walletPaymentU2URequestBody, userSender, userReceiver models.User) {
	transaction.EntityId = 0
	transaction.DateTime = time.Time{}
	transaction.TransactionMetaId = models.TransactionMeta{
		UserInitiatorId:      reqBody.UserId,
		UserReceiverID:       *userSender.Id,
		UserSenderId:         *userReceiver.Id,
		UserWalletReceiverID: reqBody.ReceiverWalletId,
		UserWalletSenderID:   reqBody.SenderWalletId,
	}
}

func (h walletHandler) WalletPaymentU2U(w http.ResponseWriter, r *http.Request) {
	var reqBody walletPaymentU2URequestBody
	body, err := ioutil.ReadAll(r.Body)

	var senderWallet, receiverWallet models.Wallet

	if err != nil {
		fmt.Fprintf(w, "err %q %q\n", err, err.Error())
		return
	}

	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		fmt.Println(w, "can't unmarshal: ", err.Error())
	}

	var newTransactionR, newTransactionS models.Transactions
	var userSender, userReceiver models.User

	senderWallet.FindWallet(reqBody.SenderWalletId, h.DB)
	receiverWallet.FindWallet(reqBody.ReceiverWalletId, h.DB)

	userSender.FindUserForWalletId(reqBody.SenderWalletId, h.DB)
	userReceiver.FindUserForWalletId(reqBody.ReceiverWalletId, h.DB)

	fillTransaction(&newTransactionR, reqBody, userSender, userReceiver)
	fillTransaction(&newTransactionS, reqBody, userSender, userReceiver)

	senderWallet.WalletTrasactions = []models.Transactions{newTransactionR}
	receiverWallet.WalletTrasactions = []models.Transactions{newTransactionS}

	resultPayment := reqBody.PaymentValue

	if reqBody.Taxes > 0 {
		taxes := resultPayment * float64(reqBody.Taxes) / 100
		resultPayment = resultPayment - taxes
		var taxesWallet models.Wallet
		taxesWallet.FindWallet(4, h.DB)
		taxesWallet.ChangeEtheriumBalance(taxes)
		h.DB.Save(&taxesWallet)
	}

	senderWallet.ChangeEtheriumBalance(-reqBody.PaymentValue)
	receiverWallet.ChangeEtheriumBalance(resultPayment)

	fmt.Println(senderWallet.EthBalance, receiverWallet.EthBalance)

	h.DB.Save(&senderWallet)
	h.DB.Save(&receiverWallet)
}
