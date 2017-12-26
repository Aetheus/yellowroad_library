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
	"yellowroad_library/database/repo/booktag_repo"
	"yellowroad_library/database/repo/booktagcount_repo"
	"yellowroad_library/database/repo/booktag_repo/gorm_booktag_repo"
	"yellowroad_library/database/repo/booktagcount_repo/gorm_booktagcount_repo"
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

func (ac AppContainer) StoryService(work uow.UnitOfWork) story_serv.StoryService {
	if (work == nil){
		work = ac.UnitOfWork()
	}
	return story_serv.Default(work);
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

func (ac AppContainer) GetBookTagRepository() booktag_repo.BookTagRepository{
	return gorm_booktag_repo.New(ac.GetDbConn())
}
func (ac AppContainer) GetBookTagCountRepository() booktagcount_repo.BookTagCountRepository{
	return gorm_booktagcount_repo.New(ac.GetDbConn())
}

func (ac AppContainer) UnitOfWork() uow.UnitOfWork {
	return uow.NewAppUnitOfWork(ac.GetDbConn())
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
