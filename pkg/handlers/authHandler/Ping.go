package authHandler

import (
	"net/http"
)

func (h authHandler) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
