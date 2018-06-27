package auth_domain

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/repo/user_repo"
	"yellowroad_library/services/token_serv"
)

type GetLoggedInUser struct {
	userRepo user_repo.UserRepository
	tokenHelper TokenStringValidator
}

func NewGetLoggedInUser(
	userRepo user_repo.UserRepository,
	tokenHelper TokenStringValidator,
) GetLoggedInUser {
	return GetLoggedInUser {userRepo,tokenHelper}
}

func (this GetLoggedInUser) Execute(tokenString string) (entities.User, app_error.AppError) {
	var user entities.User
	var err app_error.AppError

	tokenClaim, err := this.tokenHelper.ValidateTokenString(tokenString)
	if err != nil {
		return user, err
	}

	user, err = this.userRepo.FindById(tokenClaim.UserID)
	return user, err
}


type LoginClaimExtractor interface{
	GetLoginClaim()(token_serv.LoginClaim, app_error.AppError)
}
