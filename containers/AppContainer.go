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

	"github.com/jinzhu/gorm"
	"yellowroad_library/database/repo/book_repo"
	"yellowroad_library/database/repo/book_repo/gorm_book_repo"
	"yellowroad_library/services/book_serv"
	"yellowroad_library/services/story_serv"
	"yellowroad_library/database/repo/chapter_repo"
	"yellowroad_library/database/repo/chapter_repo/gorm_chapter_repo"
	"yellowroad_library/database/repo/chapterpath_repo"
	"yellowroad_library/database/repo/chapterpath_repo/gorm_chapterpath_repo"
	"yellowroad_library/database/repo/uow"
	"yellowroad_library/services/chapter_serv"
)

type AppContainer struct {
	dbConn        *gorm.DB
	tokenService  *token_serv.TokenService
	authService   *auth_serv.AuthService
	bookService	  *book_serv.BookService
	storyService  *story_serv.StoryService
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

func (ac AppContainer) AuthServiceFactory() auth_serv.AppliedAuthServiceFactory {
	return func(work uow.UnitOfWork) auth_serv.AuthService{
		return auth_serv.Default(work, ac.TokenService());
	}
}

//only needs to be resolved once, since all it needs is Configuration and that isn't dynamic,
//so no need to return a factory.
func (ac AppContainer) TokenService() token_serv.TokenService {
	if ac.tokenService == nil {
		var tokenService token_serv.TokenService = token_serv.Default(ac.GetConfiguration())
		ac.tokenService = &tokenService
	}

	return *ac.tokenService
}

func (ac AppContainer) BookServiceFactory() book_serv.BookServiceFactory {
	return book_serv.Default;
	//return func(work uow.UnitOfWork) book_serv.BookService {
	//	var bookService = book_serv.Default(work)
	//	return bookService
	//}
}

func (ac AppContainer) ChapterServiceFactory() func (work uow.UnitOfWork) chapter_serv.ChapterService {
	return chapter_serv.Default;
}

func (ac AppContainer) StoryServiceFactory() story_serv.StoryServiceFactory {
	return story_serv.Default;
	//return func(work uow.UnitOfWork) story_serv.StoryService{
	//	var storyService = story_serv.Default(work)
	//	return storyService
	//}
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

func (ac AppContainer) GetChapterRepository() chapter_repo.ChapterRepository {
	return gorm_chapter_repo.New(ac.GetDbConn())
}

func (ac AppContainer) GetChapterPathRepository() chapterpath_repo.ChapterPathRepository {
	return gorm_chapterpath_repo.New(ac.GetDbConn())
}

func (ac AppContainer) UnitOfWorkFactory() uow.SimpleUnitOfWorkFactory {
	return func() uow.UnitOfWork{
		return uow.NewAppUnitOfWork(ac.GetDbConn())
	}
}

//func (ac AppContainer) UnitOfWork() uow.UnitOfWork {
//	return uow.NewAppUnitOfWork(ac.GetDbConn())
//}

/***********************************************************************************************/
/***********************************************************************************************/
//Middleware

func (ac AppContainer) GetAuthMiddleware() auth_middleware.AuthMiddleware {
	return auth_middleware.New(ac.TokenService())
}

/***********************************************************************************************/
/***********************************************************************************************/
