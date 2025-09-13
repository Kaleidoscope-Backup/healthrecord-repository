package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/karte/microservice-utilities/auth0"
)

func AuthScopeHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ctx := r.Context()
	var (
		isAuthorized = false
	)

	authHeaderParts := strings.Split(r.Header.Get("Authorization"), " ")
	token := authHeaderParts[1]

	hasReadScope := auth0.CheckScope("read:healthrecords", token)
	hasWriteScope := auth0.CheckScope("write:healthrecords", token)

	if !hasWriteScope || !hasReadScope {
		message := "Bad Auth: Insufficient scope."
		responseJSON(message, rw, http.StatusForbidden)
		return
	}

	isAuthorized = true
	ctx = context.WithValue(ctx, "is_authorized", isAuthorized)

	next(rw, r)
}

func responseJSON(message string, w http.ResponseWriter, statusCode int) {
	response := auth0.Response{message}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}
