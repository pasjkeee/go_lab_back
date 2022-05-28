package defaultData

import (
	"awesomeProject/pkg/models"
	"awesomeProject/pkg/utils"
	"time"
)

var UserAccaunt = []models.UserAccaunt{
	{Email: "testEmail1", Name: "testName1", Surname: "testSurname1"},
	{Email: "testEmail2", Name: "testName2", Surname: "testSurname2"},
	{Email: "testEmail3", Name: "testName3", Surname: "testSurname3"},
	{Email: "taxes@gmail.com", Name: "worker", Surname: "worker"},
}

var ESignature = []models.ESingnature{
	{Signature: "testSignature1", PublicKey: "testPublicKey1"},
	{Signature: "testSignature2", PublicKey: "testPublicKey2"},
	{Signature: "testSignature3", PublicKey: "testPublicKey3"},
	{Signature: "testSignature4", PublicKey: "testPublicKey4"},
}

var AuthStatus = []models.AuthStatus{
	{Token: "testToken1", AuthLocation: "testLocation1"},
	{Token: "testToken2", AuthLocation: "testLocation1"},
	{Token: "testToken3", AuthLocation: "testLocation1"},
	{Token: "testToken4", AuthLocation: "testLocation4"},
}

var Transactions = []models.Transactions{
	{DateTime: time.Date(2021, 10, 10, 10, 30, 0, 0, time.UTC), EntityId: 1},
	{DateTime: time.Date(2021, 9, 10, 10, 30, 0, 0, time.UTC), EntityId: 1},
	{DateTime: time.Date(2021, 11, 10, 10, 30, 0, 0, time.UTC), EntityId: 1},
}

var Wallet = []models.Wallet{
	{EthBalance: 0, GasBalance: 0, WalletTrasactions: []models.Transactions{Transactions[0]}},
	{EthBalance: 100, GasBalance: 10, WalletTrasactions: []models.Transactions{Transactions[1]}},
	{EthBalance: 10, GasBalance: 0, WalletTrasactions: []models.Transactions{Transactions[2]}},
	{EthBalance: 0, GasBalance: 0, WalletTrasactions: nil},
}

var UserPermissionsGroups = []models.UserPermissionsGroups{
	{Value: "UNAUTHORIZED", Id: 0},
	{Value: "USER", Id: 1},
	{Value: "OPERATOR", Id: 2},
	{Value: "ADMIN", Id: 3},
	{Value: "NFT_OWNER", Id: 4},
	{Value: "OWNER", Id: 5},
}

var ids = []int{1, 2, 3, 4}

var pswdHashed, _ = utils.HashPassword("123")

var User = []models.User{
	{Login: "testLogin1", Password: pswdHashed,
		ESingnatureId: &ESignature[0], UserAccaunt: &UserAccaunt[0],
		AuthStatuses: &[]models.AuthStatus{AuthStatus[0]},
		Wallets:      &[]models.Wallet{Wallet[0]},
	},
	{Login: "testLogin2", Password: pswdHashed,
		ESingnatureId: &ESignature[1], UserAccaunt: &UserAccaunt[1],
		AuthStatuses: &[]models.AuthStatus{AuthStatus[1]},
		Wallets:      &[]models.Wallet{Wallet[1]},
	},
	{Login: "testLogin3", Password: pswdHashed,
		ESingnatureId: &ESignature[2], UserAccaunt: &UserAccaunt[2],
		AuthStatuses: &[]models.AuthStatus{AuthStatus[2]},
		Wallets:      &[]models.Wallet{Wallet[2]},
	},
	{Login: "taxes", Password: pswdHashed,
		ESingnatureId: &ESignature[3], UserAccaunt: &UserAccaunt[3],
		AuthStatuses: &[]models.AuthStatus{AuthStatus[3]},
		Wallets:      &[]models.Wallet{Wallet[3]},
	},
}

var Nft = []models.Nft{
	{MetaId: 1, LocationLink: "location1", OwnerId: 1, CreatorId: 1},
	{MetaId: 2, LocationLink: "location2", OwnerId: 2, CreatorId: 1},
	{MetaId: 3, LocationLink: "location3", OwnerId: 3, CreatorId: 1},
}

var BalanceHistory = []models.BalanceHistory{
	{EthBalance: 0, GasBalance: 0, WalletId: 1},
	{EthBalance: 0, GasBalance: 0, WalletId: 2},
	{EthBalance: 0, GasBalance: 0, WalletId: 3},
	{EthBalance: 100, GasBalance: 2, WalletId: 1},
	{EthBalance: 2000, GasBalance: 1, WalletId: 2},
	{EthBalance: 10, GasBalance: 0, WalletId: 3},
}

var TransactionMeta = []models.TransactionMeta{
	{UserInitiatorId: 1, UserReceiverID: 2, UserSenderId: 3, UserWalletReceiverID: 2, UserWalletSenderID: 3, UserReceiverPublicSignature: "testPublicKey2", UserSenderPublicSignature: "testPublicKey3", TransactionId: 1},
	{UserInitiatorId: 2, UserReceiverID: 1, UserSenderId: 2, UserWalletReceiverID: 1, UserWalletSenderID: 2, UserReceiverPublicSignature: "testPublicKey1", UserSenderPublicSignature: "testPublicKey2", TransactionId: 2},
	{UserInitiatorId: 3, UserReceiverID: 1, UserSenderId: 1, UserWalletReceiverID: 1, UserWalletSenderID: 1, UserReceiverPublicSignature: "testPublicKey1", UserSenderPublicSignature: "testPublicKey1", TransactionId: 3},
}
