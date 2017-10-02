package app_token_serv

import (
	"testing"
	"yellowroad_library/database/entities"
)

var user = 	entities.User{
				ID : 95,
				Username : "bob",
				Password : "az89712,lazlkjap109asnklawjlj34",
				Email    : "bob@mail.com",
			}


func TestCreateTokenString(t *testing.T) {
	var appTokenServ AppTokenService = New()
	tokenString, err := appTokenServ.CreateTokenString(user)
	if err != nil {
		t.Errorf("Error occured:\n -Context Message: %s\n - EndpointMessage: %s", err.Error(), err.EndpointMessage())
	}

	if len(tokenString) == 0 {
		t.Errorf("Token string should not have a lenght of zero!")
	}

}

func TestValidateTokenString(t *testing.T) {
	var appTokenServ AppTokenService = New()
	tokenString, _ := appTokenServ.CreateTokenString(user)

	claim, err := appTokenServ.ValidateTokenString(tokenString)
	if err != nil {
		t.Errorf("Error occured:\n -Context Message: %s\n - EndpointMessage: %s",
			err.Error(), err.EndpointMessage())
	}

	if claim.UserID != user.ID {
		t.Errorf("The claim's user id of [%d] does not match the expected user ID of [%d]!",
			claim.UserID, user.ID)
	}
}