package auth_domain

import (
	"yellowroad_library/database/entities"
	"testing"
	"github.com/dgrijalva/jwt-go"
)

var user = 	entities.User{
	ID : 95,
	Username : "bob",
	Password : "az89712,lazlkjap109asnklawjlj34",
	Email    : "bob@mail.com",
}


func TestCreateTokenString(t *testing.T) {
	tokenHelper := NewTokenHelper(jwt.SigningMethodHS256,[]byte("verySecretKey"), 10)

	tokenString, err := tokenHelper.CreateTokenString(user)
	if err != nil {
		t.Errorf("Error occured:\n -Context Message: %s\n - EndpointMessage: %s", err.Error(), err.EndpointMessage())
	}

	if len(tokenString) == 0 {
		t.Errorf("Token string should not have a lenght of zero!")
	}
}

func TestValidateTokenString(t *testing.T) {
	tokenHelper := NewTokenHelper(jwt.SigningMethodHS256,[]byte("verySecretKey"), 10)
	tokenString, _ := tokenHelper.CreateTokenString(user)

	claim, err := tokenHelper.ValidateTokenString(tokenString)
	if err != nil {
		t.Errorf("Error occured:\n -Context Message: %s\n - EndpointMessage: %s",
			err.Error(), err.EndpointMessage())
	}

	if claim.UserID != user.ID {
		t.Errorf("The claim's user id of [%d] does not match the expected user ID of [%d]!",
			claim.UserID, user.ID)
	}
}