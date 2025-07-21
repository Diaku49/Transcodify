package utilities

import (
	"fmt"
	"net/http"

	"github.com/Diaku49/FoodOrderSystem/backend/internals/constants"
	"github.com/golang-jwt/jwt/v4"
)

func GetUserIDFromContext(r *http.Request) (uint, error) {
	claimsVal := r.Context().Value(constants.ClaimsKey)
	claims, ok := claimsVal.(jwt.MapClaims)
	if !ok || claimsVal == nil {
		return 0, fmt.Errorf("claims missing or invalid")
	}

	userIdVal, ok := claims["user_id"]
	userIdFloat, ok2 := userIdVal.(float64)
	if !ok || !ok2 {
		return 0, fmt.Errorf("id missing or invalid")
	}

	return uint(userIdFloat), nil
}
