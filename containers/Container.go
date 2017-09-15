package containers

import (
	"yellowroad_library/config"
	"yellowroad_library/database/repo/user_repo"
	"yellowroad_library/http/middleware/auth_middleware"
	"yellowroad_library/services/auth_serv"
	"yellowroad_library/services/token_serv"
)

/*Container :
Do NOT modify the Container interface to include generating any low-level database
related structs (e.g: creating a GORM database connection). Remember, Repositories
are our abstraction over database access, and the Repository interfaces themselves are
DB driver agnostic (even if their implementations are of course not).

The _implementations_ of Container can, of course, generate their own DB connections
and pass these to the Repositories (and will almost definitely need to do so)
*/
type Container interface {
	//services
	GetAuthService() auth_serv.AuthService
	GetTokenService() token_serv.TokenService

	//repo
	GetUserRepository() user_repo.UserRepository

	//middleware
	GetAuthMiddleware() auth_middleware.AuthMiddleware

	//configuration
	GetConfiguration() config.Configuration
}
