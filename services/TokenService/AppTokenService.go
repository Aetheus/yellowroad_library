package TokenService

import (
	"fmt"
	"time"

	"yellowroad_library/database/entities"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

const TOKEN_CLAIMS_CONTEXT_KEY = "TokenClaims"

type AppTokenService struct {
	dbConn               *gorm.DB
	signingMethod        jwt.SigningMethod
	secretKey            []byte
	expiryDurationInDays int
}

func NewAppTokenService(dbConn *gorm.DB) AppTokenService {
	return AppTokenService{
		dbConn:               dbConn,
		signingMethod:        jwt.SigningMethodHS256, //default this for now
		secretKey:            []byte("blubber"),      //default this for now
		expiryDurationInDays: 365,                    //default this for now
	}
}

func (service AppTokenService) ValidateTokenString(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return service.secretKey, nil
	})

	fmt.Println(33)
	if err != nil {
		return nil, err
	}

	fmt.Println(38)
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		fmt.Println(claims)
		return claims, nil
	}

	return nil, nil
}

func (service AppTokenService) CreateTokenString(user entities.User) (string, error) {
	nowTimestamp := time.Now().Unix()
	expiryDate := time.Now().AddDate(0, 0, service.expiryDurationInDays).Unix()

	claims := MyCustomClaims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: expiryDate,
			NotBefore: nowTimestamp,
			IssuedAt:  nowTimestamp,
			Issuer:    "yellowroad",
		},
	}

	token := jwt.NewWithClaims(
		service.signingMethod,
		claims,
	)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(service.secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
