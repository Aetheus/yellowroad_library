package auth_serv

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/repo/uow"
	"yellowroad_library/services/token_serv"
)

type AuthService interface {
	RegisterUser(username string, password string, email string) (entities.User, app_error.AppError)
	LoginUser(username string, password string) (entities.User, string, app_error.AppError)
	VerifyToken(token string) (entities.User, app_error.AppError)
	GetLoggedInUser(adapter LoginClaimExtractor) (entities.User, app_error.AppError)

	SetUnitOfWork(work uow.UnitOfWork)
}

type LoginClaimExtractor interface{
	GetLoginClaim()(token_serv.LoginClaim, app_error.AppError)
}