package main

import (
	"awesomeProject/pkg/config"
	"awesomeProject/pkg/db"
	"awesomeProject/pkg/handlers/authHandler"
	"awesomeProject/pkg/handlers/userHandler"
	"awesomeProject/pkg/handlers/walletHandler"
	"awesomeProject/pkg/logging"
	"awesomeProject/pkg/middleware"
	"awesomeProject/pkg/socket"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {

	//if len(os.Args) < 4 {
	//	panic("usage go run main.go [dbUser] [dbPswd] [dbName]")
	//}

	logging.Init()
	logger := logging.GetLogger()
	logger.Println("logger initialized")

	logger.Println("config initializing")
	cfg := config.GetConfig()

	DB := db.Init(cfg)

	router := mux.NewRouter()
	userH := userHandler.New(DB)
	walletH := walletHandler.New(DB)
	loginH := authHandler.New(DB)

	router.HandleFunc("/login", loginH.SignIn).Methods(http.MethodPost)
	router.HandleFunc("/logout", loginH.SignOut).Methods(http.MethodPost)
	router.HandleFunc("/signup", loginH.SignUp).Methods(http.MethodPost)
	router.HandleFunc("/ping", loginH.Ping).Methods(http.MethodGet)

	router.HandleFunc("/users", userH.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{login}", userH.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/users", userH.AddUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{login}", userH.UpdatePasswordUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{login}", userH.DeleteUser).Methods(http.MethodDelete)

	router.HandleFunc("/wallet/{login}", walletH.AddWallet).Methods(http.MethodPost)
	router.HandleFunc("/wallet/payment/u2u", walletH.WalletPaymentU2U).Methods(http.MethodPost)
	router.HandleFunc("/wallet/{id}", walletH.GetUserWallets).Methods(http.MethodGet)
	router.HandleFunc("/wallet/transactions/{id}", walletH.GetWalletTransactions).Methods(http.MethodGet)

	router.HandleFunc("/ws/start", socket.Echo)

	router.Use(middleware.CheckAuthMiddleware)

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "UPGRADE"},
		AllowedHeaders: []string{"X-Requested-With", "Content-Type", "Authorization", "Access-Control-Allow-Origin", "Content-Type",
			"Connection", "Host", "Sec-WebSocket-Key", "Sec-WebSocket-Version", "Sec-WebSocket-Extensions"},
		MaxAge:               86400,
		AllowCredentials:     true,
		OptionsSuccessStatus: 204,
		Debug:                true,
	}).Handler(router)
	log.Println("API is running!")

	log.Fatal(http.ListenAndServe(":8080", handler))
}
