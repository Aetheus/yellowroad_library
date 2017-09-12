package tokenService

import (
	"yellowroad_library/database/entities"

	jwt "github.com/dgrijalva/jwt-go"
)

const TOKEN_CLAIMS_CONTEXT_KEY = "TokenClaims"

type TokenService interface {
	ValidateTokenString(tokenString string) (*MyCustomClaims, error)
	CreateTokenString(user entities.User) (string, error)
}

type MyCustomClaims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}
