package token_serv

import (
	"time"

	"yellowroad_library/database/entities"

	"github.com/dgrijalva/jwt-go"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/config"
	"fmt"
)

type DefaultTokenService struct {
	signingMethod        jwt.SigningMethod
	secretKey            []byte
	expiryDurationInDays int
}
//ensure interface implementation
var _ TokenService = DefaultTokenService{}

func Default(configuration config.Configuration) DefaultTokenService {
	return DefaultTokenService{
		signingMethod:        jwt.SigningMethodHS256,
		secretKey:            []byte(configuration.JWT.SecretKey),
		expiryDurationInDays: configuration.JWT.ExpiryDurationInDays,
	}
}

func (service DefaultTokenService) ValidateTokenString(tokenString string) (LoginClaim, app_error.AppError) {
	var claims LoginClaim

	token, err := jwt.ParseWithClaims(tokenString, &LoginClaim{}, func(token *jwt.Token) (interface{}, error) {

		if token.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}


		return service.secretKey, nil
	})



	if err != nil {
		return claims, app_error.Wrap(err)
	}

	if claims, ok := token.Claims.(*LoginClaim); ok && token.Valid {
		return *claims, nil
	}

	return claims, nil
}

func (service DefaultTokenService) CreateTokenString(user entities.User) (string, app_error.AppError) {
	nowTimestamp := time.Now().Unix()
	expiryDate := time.Now().AddDate(0, 0, service.expiryDurationInDays).Unix()

	claims := LoginClaim{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: expiryDate,
			NotBefore: nowTimestamp,
			IssuedAt:  nowTimestamp,
			Issuer:    "yellowroad",	//TODO: make the issuer configurable or at least extract it to a constant
		},
	}

	token := jwt.NewWithClaims(
		service.signingMethod,
		claims,
	)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(service.secretKey)

	if err != nil {
		return "", app_error.Wrap(err)
	}

	return tokenString, nil
}
