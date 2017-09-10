package containers

import (
	"yellowroad_library/configs"
	"yellowroad_library/database/repositories/UserRepo"
	"yellowroad_library/http/middleware/AuthMiddleware"
	"yellowroad_library/services/AuthService"
	"yellowroad_library/services/TokenService"
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
	GetAuthService() AuthService.AuthService
	GetTokenService() TokenService.TokenService

	//repositories
	GetUserRepository() UserRepo.UserRepository

	//middleware
	GetAuthMiddleware() AuthMiddleware.AuthMiddleware

	//configuration
	GetConfiguration() configs.Configuration
}
