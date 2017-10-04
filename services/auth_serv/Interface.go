package auth_serv

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
)

type AuthService interface {

	LoginUser(username string, password string) (entities.User, string, app_error.AppError)

	//TODO: unfortunately, unless we want a tonne of boilerplate, this is the "quickest" data
	//from the routes - by actually passing the context object to it. Now why are we passing
	//it as an empty interface instead of as a *gin.Context? In case we ditch Gin, mostly. 
	//may not be a good idea, so may have to revisit this
	GetLoggedInUser(data interface{}) (entities.User, app_error.AppError)
}
