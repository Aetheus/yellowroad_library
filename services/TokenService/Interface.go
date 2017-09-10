package TokenService

import (
	"yellowroad_library/database/entities"

	jwt "github.com/dgrijalva/jwt-go"
)

type TokenService interface {
	ValidateTokenString(tokenString string) (*MyCustomClaims, error)
	CreateTokenString(user entities.User) (string, error)
}

type MyCustomClaims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}
