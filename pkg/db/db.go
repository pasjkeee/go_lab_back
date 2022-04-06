package db

import (
	"awesomeProject/pkg/defaultData"
	"awesomeProject/pkg/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func InsertInto(db *gorm.DB, userId, permId int) {
	db.Exec("INSERT INTO user_permissions (user_id, user_permissions_groups_id) VALUES ($1, $2)", userId, permId)
}

func Init() *gorm.DB {
	//dbURL := "postgres://postgres:pavel@localhost:5432/nft"
	dbURL := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", os.Args[1], os.Args[2], os.Args[3])

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	db.Exec("drop table nfts;\n\ndrop table user_accaunts;\n\ndrop table wallet_transactions;\n\ndrop table user_wallets;\n\ndrop table e_singnatures;\n\ndrop table user_auth_status;\n\ndrop table auth_statuses;\n\ndrop table balance_histories;\n\ndrop table wallets;\n\ndrop table transaction_meta;\n\ndrop table transactions;\n\ndrop table user_permissions;\n\ndrop table users;\n\ndrop table user_permissions_groups;\n\n")

	if err != nil {
		log.Fatalln(err)
	}

	e := db.AutoMigrate(
		&models.UserPermissionsGroups{},
		&models.User{},
		&models.Nft{},
		&models.UserAccaunt{},
		&models.Wallet{},
		&models.ESingnature{},
		&models.AuthStatus{},
		&models.BalanceHistory{},
		&models.Transactions{},
		&models.TransactionMeta{},
		&models.RefreshToken{},
	)

	if err != nil {
		log.Fatalln(e)
		return db
	}

	for _, data := range defaultData.Transactions {
		db.Create(&data)
	}

	for _, data := range defaultData.TransactionMeta {
		db.Create(&data)
	}

	for _, data := range defaultData.User {
		db.Create(&data)
	}

	for _, data := range defaultData.UserPermissionsGroups {
		db.Create(&data)
	}

	for _, data := range defaultData.Nft {
		db.Create(&data)
	}

	for _, data := range defaultData.BalanceHistory {
		db.Create(&data)
	}

	InsertInto(db, 1, 1)
	InsertInto(db, 2, 1)
	InsertInto(db, 3, 1)

	return db
}
