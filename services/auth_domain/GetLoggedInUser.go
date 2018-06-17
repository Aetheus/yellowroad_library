package auth_domain

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/repo/user_repo"
	"yellowroad_library/services/token_serv"
)

type GetLoggedInUser struct {
	userRepo user_repo.UserRepository
}

func (this GetLoggedInUser) Execute(loginClaimAdapter LoginClaimExtractor) (entities.User, app_error.AppError) {
	var user entities.User
	var err app_error.AppError

	tokenClaim, err := loginClaimAdapter.GetLoginClaim()
	if err != nil {
		return user, app_error.Wrap(err)
	}

	user, err = this.userRepo.FindById(tokenClaim.UserID)
	return user, err
}


type LoginClaimExtractor interface{
	GetLoginClaim()(token_serv.LoginClaim, app_error.AppError)
}
