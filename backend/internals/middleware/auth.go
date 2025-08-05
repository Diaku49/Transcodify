package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	Jwt "github.com/Diaku49/FoodOrderSystem/backend/internals/JwtService"
	"github.com/Diaku49/FoodOrderSystem/backend/internals/constants"
	util "github.com/Diaku49/FoodOrderSystem/backend/utilities"
)

func Auth(next http.HandlerFunc, secret []byte) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			util.WriteJsonError(w, "authorization failed", http.StatusForbidden, errors.New("no token provided"))
			return
		}
		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

		claims, err := Jwt.ParseJwt(secret, tokenStr)
		if err != nil {
			util.WriteJsonError(w, "authorization failed", http.StatusForbidden, err)
			return
		}

		// set claims
		ctx := context.WithValue(r.Context(), constants.ClaimsKey, claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
