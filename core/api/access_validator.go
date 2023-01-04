package api

import (
	"net/http"
	"strings"

	"github.com/lekoller/dokumentar/core/shared"
)

func validateAccess(w http.ResponseWriter, r *http.Request) (isValid bool) {
	config, err := shared.LoadConfig(".")
	if err != nil {
		return
	}
	auth := r.Header.Get("Authorization")

	if len(strings.Split(auth, "Bearer ")) == 2 {
		token := strings.Split(auth, "Bearer ")[1]
		isValid = token == config.APIToken
	}
	if !isValid {
		w.WriteHeader(401)
	}
	return
}
