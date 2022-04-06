package models

import "gorm.io/gorm"

type TransactionMeta struct {
	gorm.Model
	Id                          int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	UserInitiatorId             int    `json:"user_initiator_id"`
	UserReceiverID              int    `json:"user_receiver_id"`
	UserSenderId                int    `json:"user_sender_id"`
	UserWalletReceiverID        int    `json:"user_wallet_receiver_id"`
	UserWalletSenderID          int    `json:"user_wallet_sender_id"`
	UserReceiverPublicSignature string `json:"user_receiver_public_signature"`
	UserSenderPublicSignature   string `json:"user_sender_public_signature"`
	TransactionId               int    `json:"transaction_id"`
}
