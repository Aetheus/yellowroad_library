package token_serv

import (
	"yellowroad_library/database/entities"

	jwt "github.com/dgrijalva/jwt-go"
	"yellowroad_library/utils/app_error"
)

const TOKEN_CLAIMS_CONTEXT_KEY = "TokenClaims"

type TokenService interface {
	ValidateTokenString(tokenString string) (LoginClaim, app_error.AppError)
	CreateTokenString(user entities.User) (string, app_error.AppError)
}

type LoginClaim struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}
