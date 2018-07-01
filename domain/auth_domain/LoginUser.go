package auth_domain

import (
	"yellowroad_library/database/repo/user_repo"
	"yellowroad_library/database/entities"
	"net/http"
	"yellowroad_library/utils/app_error"
	"golang.org/x/crypto/bcrypt"
)

type LoginUser struct {
	userRepo user_repo.UserRepository
	tokenCreator TokenStringCreator
}

func NewLoginUser(userRepo user_repo.UserRepository, tokenCreator TokenStringCreator) LoginUser{
	return LoginUser{userRepo, tokenCreator}
}

func (this LoginUser) Execute(username string, password string) (entities.User, string, app_error.AppError){
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