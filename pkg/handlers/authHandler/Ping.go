package authHandler

import (
	"awesomeProject/pkg/logging"
	"net/http"
)

func (h authHandler) Ping(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()
	logger.Info("Ping")
	w.WriteHeader(http.StatusOK)

}
