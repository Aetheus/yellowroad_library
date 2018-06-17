package auth_domain

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"yellowroad_library/database/repo/user_repo"
)

type TokenStringCreator interface {
	CreateTokenString(entities.User) (string, app_error.AppError)
}

type RegisterUser struct {
	userRepo user_repo.UserRepository
	tokenCreator TokenStringCreator
}

func NewRegisterUser(userRepo user_repo.UserRepository, tokenCreator TokenStringCreator) RegisterUser{
	return RegisterUser{userRepo, tokenCreator}
}

func (this RegisterUser) Execute(username string, password string, email string) (entities.User, string, app_error.AppError){
	var user entities.User
	var err error

	//TODO : email as well
	user, findErr := this.userRepo.FindByUsername(username)
	if findErr != nil {
		if (findErr.HttpCode() == http.StatusNotFound){
			findErr = app_error.New(http.StatusUnauthorized, "","Incorrect username or password")
		}
		return user, "", findErr
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return user, "", app_error.New(http.StatusUnauthorized, "","Incorrect username or password")
	}

	token, err := this.tokenCreator.CreateTokenString(user)
	if err != nil {
		return user, "", app_error.Wrap(err)
	}

	return user, token, nil
}