package auth_serv

import (
	"yellowroad_library/database/entities"
)

type AuthService interface {
	RegisterUser(username string, password string, email string) (*entities.User, error)
	LoginUser(username string, password string) (*entities.User, string, error)

	//TODO: unfortunately, unless we want a tonne of boilerplate, this is the "quickest" data
	//from the routes - by actually passing the context object to it. Now why are we passing
	//it as an empty interface instead of as a *gin.Context? In case we ditch Gin, mostly. 
	//may not be a good idea, so may have to revisit this
	GetLoggedInUser(data interface{}) (*entities.User, error)
}
