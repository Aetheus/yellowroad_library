package containers

import (
	"fmt"

	"yellowroad_library/config"
	"yellowroad_library/http/middleware/auth_middleware"
	"yellowroad_library/services/auth_serv"
	"yellowroad_library/services/token_serv"

	"github.com/jinzhu/gorm"
	"yellowroad_library/services/book_serv"
	"yellowroad_library/services/game_serv"
	"yellowroad_library/database/repo/uow"
	"yellowroad_library/services/chapter_serv"
	"yellowroad_library/utils/app_error"
)

type AppContainer struct {
	dbConn        *gorm.DB
	tokenService  *token_serv.TokenService
	authService   *auth_serv.AuthService
	bookService	  *book_serv.BookService
	storyService  *game_serv.GameService
	configuration config.Configuration
}
//ensure interface implementation
var _ Container = AppContainer{}

func NewAppContainer(config config.Configuration) (container AppContainer, appErr app_error.AppError) {
	//setup the DB connection
	var dbSettings = config.Database
	var dbType = dbSettings.Driver
	var connectionString = fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=%s password=%s",
		dbSettings.Host, dbSettings.Username, dbSettings.Database, dbSettings.SSLMode, dbSettings.Password,
	)
	dbConn, err := gorm.Open(dbType, connectionString)
	dbConn.LogMode(true)	//TODO : only do this if debug mode enabled
	if err != nil {
		appErr = app_error.Wrap(err)
		return
	}


	container = AppContainer {
		configuration: config,
		dbConn: dbConn,
	}
	return
}

/***********************************************************************************************/
/***********************************************************************************************/
//Non-interface methods
//


/***********************************************************************************************/
/***********************************************************************************************/
//Configuration

func (ac AppContainer) GetConfiguration() config.Configuration {
	return ac.configuration
}

/***********************************************************************************************/
/***********************************************************************************************/
//Services

func (ac AppContainer) TokenService() token_serv.TokenService {
	if ac.tokenService == nil {
		var tokenService token_serv.TokenService = token_serv.Default(ac.GetConfiguration())
		ac.tokenService = &tokenService
	}

	return *ac.tokenService
}

func (ac AppContainer) AuthService(work uow.UnitOfWork) auth_serv.AuthService{
	if (work == nil){
		work = ac.UnitOfWork()
	}
	return auth_serv.Default(work, ac.TokenService());
}

func (ac AppContainer) BookService(work uow.UnitOfWork) book_serv.BookService {
	if (work == nil){
		work = ac.UnitOfWork()
	}
	return book_serv.Default(work)
}

func (ac AppContainer) ChapterService(work uow.UnitOfWork) chapter_serv.ChapterService {
	if (work == nil){
		work = ac.UnitOfWork()
	}
	return chapter_serv.Default(work);
}

func (ac AppContainer) StoryService(work uow.UnitOfWork) game_serv.GameService {
	if (work == nil){
		work = ac.UnitOfWork()
	}
	return game_serv.Default(work);
}

/***********************************************************************************************/
/***********************************************************************************************/
// Unit of Work

func (ac AppContainer) UnitOfWork() uow.UnitOfWork {
	return uow.NewAppUnitOfWork(ac.dbConn)
}


/***********************************************************************************************/
/***********************************************************************************************/
//Middleware

func (ac AppContainer) GetAuthMiddleware() auth_middleware.AuthMiddleware {
	return auth_middleware.New(ac.TokenService())
}

/***********************************************************************************************/
/***********************************************************************************************/
