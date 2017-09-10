package containers

import (
	"fmt"
	"yellowroad_library/configs"
	db "yellowroad_library/database"
	"yellowroad_library/database/repositories/UserRepo"
	"yellowroad_library/http/middleware/AuthMiddleware"
	"yellowroad_library/services/AuthService"
	"yellowroad_library/services/TokenService"

	"github.com/jinzhu/gorm"
)

type AppContainer struct {
	dbConn        *gorm.DB
	tokenService  *TokenService.TokenService
	AuthService   *AuthService.AuthService
	configuration configs.Configuration
}

func NewAppContainer(config configs.Configuration) AppContainer {
	return AppContainer{
		configuration: config,
	}
}

/***********************************************************************************************/
/***********************************************************************************************/
//Non-interface methods

func (ac AppContainer) GetDbConn() *gorm.DB {
	var dbSettings = ac.configuration.Database

	var dbType = dbSettings.Driver
	var connectionString = fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=%s password=%s",
		dbSettings.Host,
		dbSettings.Username,
		dbSettings.Database,
		dbSettings.SSLMode,
		dbSettings.Password,
	)

	if ac.dbConn == nil {
		ac.dbConn = db.Conn(dbType, connectionString)
	}

	return ac.dbConn
}

/***********************************************************************************************/
/***********************************************************************************************/
//Configuration

func (ac AppContainer) GetConfiguration() configs.Configuration {
	return ac.configuration
}

/***********************************************************************************************/
/***********************************************************************************************/
//Services

func (ac AppContainer) GetAuthService() AuthService.AuthService {
	if ac.AuthService == nil {
		var AuthService AuthService.AuthService = AuthService.NewAppAuthService(ac.GetUserRepository(), ac.GetTokenService())
		ac.AuthService = &AuthService
	}

	return *ac.AuthService
}

func (ac AppContainer) GetTokenService() TokenService.TokenService {
	if ac.tokenService == nil {
		var tokenService TokenService.TokenService = TokenService.NewAppTokenService(ac.GetDbConn())
		ac.tokenService = &tokenService
	}

	return *ac.tokenService
}

/***********************************************************************************************/
/***********************************************************************************************/
//Repositories

func (ac AppContainer) GetUserRepository() UserRepo.UserRepository {
	return UserRepo.NewGormSqlUserRepository(ac.GetDbConn())
}

/***********************************************************************************************/
/***********************************************************************************************/
//Middleware

func (ac AppContainer) GetAuthMiddleware() AuthMiddleware.AuthMiddleware {
	return AuthMiddleware.New(ac.GetTokenService())
}

/***********************************************************************************************/
/***********************************************************************************************/
