package auth_serv

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/repo/uow"
	"yellowroad_library/services/token_serv"
)

type AuthService interface {
	RegisterUser(username string, password string, email string) (*entities.User, app_error.AppError)
	LoginUser(username string, password string) (entities.User, string, app_error.AppError)

	//TODO: unfortunately, unless we want a tonne of boilerplate, this is the "quickest" data
	//from the routes - by actually passing the context object to it. Now why are we passing
	//it as an empty interface instead of as a *gin.Context? In case we ditch Gin, mostly. 
	//may not be a good idea, so may have to revisit this
	GetLoggedInUser(data interface{}) (entities.User, app_error.AppError)
}


type AuthServiceFactory func (work uow.UnitOfWork, tokenService token_serv.TokenService) AuthService

//since TokenService is a dependency that only needs to be resolved once,
//we can have a factory that only requires the truly runtime parameter (UnitOfWork)
type AppliedAuthServiceFactory func (work uow.UnitOfWork) AuthService