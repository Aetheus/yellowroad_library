package auth_domain

import (
	"yellowroad_library/utils/app_error"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"yellowroad_library/database/entities"
	"time"
)

type TokenStringValidator interface{
	ValidateTokenString(string) (LoginClaim, app_error.AppError)
}
type TokenStringCreator interface {
	CreateTokenString(entities.User) (string, app_error.AppError)
}
type TokenHelper struct {
	signingMethod        jwt.SigningMethod
	secretKey            []byte
	expiryDurationInDays int
}
var _ TokenStringValidator = TokenHelper{}
var _ TokenStringCreator = TokenHelper{}


func NewTokenHelper(
	signingMethod        jwt.SigningMethod,
	secretKey            []byte,
	expiryDurationInDays int,
) TokenHelper {
	return TokenHelper {
		signingMethod,secretKey,expiryDurationInDays,
	}
}

func (this TokenHelper) ValidateTokenString(tokenString string) (LoginClaim, app_error.AppError) {
	var claims LoginClaim

	token, err := jwt.ParseWithClaims(tokenString, &LoginClaim{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != this.signingMethod.Alg() {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return this.secretKey, nil
	})

	if err != nil {
		return claims, app_error.Wrap(err)
	}

	if claims, ok := token.Claims.(*LoginClaim); ok && token.Valid {
		return *claims, nil
	}

	return claims, nil
}

func (this TokenHelper) CreateTokenString(user entities.User) (string, app_error.AppError) {
	nowTimestamp := time.Now().Unix()
	expiryDate := time.Now().AddDate(0, 0, this.expiryDurationInDays).Unix()

	claims := LoginClaim{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: expiryDate,
			NotBefore: nowTimestamp,
			IssuedAt:  nowTimestamp,
			Issuer:    "yellowroad",	//TODO: make the issuer configurable or at least extract it to a constant
		},
	}

	token := jwt.NewWithClaims(this.signingMethod,claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(this.secretKey)

	if err != nil {
		return "", app_error.Wrap(err)
	}

	return tokenString, nil
}
