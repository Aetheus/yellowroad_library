package containers

import (
	"fmt"

	"yellowroad_library/config"
	db "yellowroad_library/database"
	"yellowroad_library/database/repo/user_repo"
	"yellowroad_library/http/middleware/auth_middleware"
	"yellowroad_library/services/auth_serv"
	"yellowroad_library/services/token_serv"
	"yellowroad_library/database/repo/user_repo/gorm_user_repo"
	"yellowroad_library/services/auth_serv/app_auth_serv"
	"yellowroad_library/services/token_serv/app_token_serv"

	"github.com/jinzhu/gorm"
	"yellowroad_library/database/repo/book_repo"
	"yellowroad_library/database/repo/book_repo/gorm_book_repo"
	"yellowroad_library/services/book_serv"
	"yellowroad_library/services/book_serv/app_book_serv"
)

type AppContainer struct {
	dbConn        *gorm.DB
	tokenService  *token_serv.TokenService
	authService   *auth_serv.AuthService
	bookService	  *book_serv.BookService
	configuration config.Configuration
}
//ensure interface implementation
var _ Container = AppContainer{}

func NewAppContainer(config config.Configuration) AppContainer {
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

func (ac AppContainer) GetConfiguration() config.Configuration {
	return ac.configuration
}

/***********************************************************************************************/
/***********************************************************************************************/
//Services

func (ac AppContainer) GetAuthService() auth_serv.AuthService {
	if ac.authService == nil {
		var AuthService auth_serv.AuthService = app_auth_serv.New(ac.GetUserRepository(), ac.GetTokenService())
		ac.authService = &AuthService
	}

	return *ac.authService
}

func (ac AppContainer) GetTokenService() token_serv.TokenService {
	if ac.tokenService == nil {
		var tokenService token_serv.TokenService = app_token_serv.New(ac.GetDbConn())
		ac.tokenService = &tokenService
	}

	return *ac.tokenService
}

func (ac AppContainer) GetBookService() book_serv.BookService {
	if ac.bookService == nil {
		var bookService book_serv.BookService = app_book_serv.New(ac.GetBookRepository(),ac.GetUserRepository())
		ac.bookService = &bookService
	}

	return *ac.bookService
}

/***********************************************************************************************/
/***********************************************************************************************/
//Repositories

func (ac AppContainer) GetUserRepository() user_repo.UserRepository {
	return gorm_user_repo.New(ac.GetDbConn())
}

func (ac AppContainer) GetBookRepository() book_repo.BookRepository {
	return gorm_book_repo.New(ac.GetDbConn())
}

/***********************************************************************************************/
/***********************************************************************************************/
//Middleware

func (ac AppContainer) GetAuthMiddleware() auth_middleware.AuthMiddleware {
	return auth_middleware.New(ac.GetTokenService())
}

/***********************************************************************************************/
/***********************************************************************************************/
