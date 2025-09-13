package auth0

import (
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	Scope string `json:"scope"`
	jwt.StandardClaims
}

func CheckScope(scope string, tokenString string) bool {
	token, _ := jwt.ParseWithClaims(tokenString, &CustomClaims{}, nil)

	claims, _ := token.Claims.(*CustomClaims)

	hasScope := false
	result := strings.Split(claims.Scope, " ")
	for i := range result {
		if result[i] == scope {
			hasScope = true
		}
	}

	return hasScope
}
