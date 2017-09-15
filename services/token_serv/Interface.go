package token_serv

import (
	"yellowroad_library/database/entities"

	jwt "github.com/dgrijalva/jwt-go"
)

const TOKEN_CLAIMS_CONTEXT_KEY = "TokenClaims"

type TokenService interface {
	ValidateTokenString(tokenString string) (LoginClaim, error)
	CreateTokenString(user entities.User) (string, error)
}

type LoginClaim struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}
