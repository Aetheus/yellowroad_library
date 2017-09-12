package containers

import (
	"fmt"

	"yellowroad_library/configs"
	db "yellowroad_library/database"
	"yellowroad_library/database/repositories/userRepository"
	"yellowroad_library/http/middleware/authMiddleware"
	"yellowroad_library/services/authService"
	"yellowroad_library/services/tokenService"
	"yellowroad_library/database/repositories/userRepository/gormUserRepository"
	"yellowroad_library/services/authService/appAuthService"
	"yellowroad_library/services/tokenService/appTokenService"

	"github.com/jinzhu/gorm"
)

type AppContainer struct {
	dbConn        *gorm.DB
	tokenService  *tokenService.TokenService
	AuthService   *authService.AuthService
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

func (ac AppContainer) GetAuthService() authService.AuthService {
	if ac.AuthService == nil {
		var AuthService authService.AuthService = appAuthService.New(ac.GetUserRepository(), ac.GetTokenService())
		ac.AuthService = &AuthService
	}

	return *ac.AuthService
}

func (ac AppContainer) GetTokenService() tokenService.TokenService {
	if ac.tokenService == nil {
		var tokenService tokenService.TokenService = appTokenService.New(ac.GetDbConn())
		ac.tokenService = &tokenService
	}

	return *ac.tokenService
}

/***********************************************************************************************/
/***********************************************************************************************/
//Repositories

func (ac AppContainer) GetUserRepository() userRepository.UserRepository {
	return gormUserRepository.New(ac.GetDbConn())
}

/***********************************************************************************************/
/***********************************************************************************************/
//Middleware

func (ac AppContainer) GetAuthMiddleware() authMiddleware.AuthMiddleware {
	return authMiddleware.New(ac.GetTokenService())
}

/***********************************************************************************************/
/***********************************************************************************************/
