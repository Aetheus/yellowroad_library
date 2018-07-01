package containers

import (
	"fmt"

	"yellowroad_library/config"
	"yellowroad_library/http/middleware/auth_middleware"
	"github.com/jinzhu/gorm"
	"yellowroad_library/database/repo/uow"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/services/auth_domain"
	"github.com/dgrijalva/jwt-go"
	"yellowroad_library/database/repo/user_repo"
	"yellowroad_library/services/book_domain"
	"yellowroad_library/services/chapter_domain"
	"yellowroad_library/services/game_domain"
)

type AppContainer struct {
	dbConn        *gorm.DB
	configuration config.Configuration
	tokenHelper   	*auth_domain.TokenHelper
}

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

func (ac AppContainer) Configuration() config.Configuration {
	return ac.configuration
}


/***********************************************************************************************/
/***********************************************************************************************/
// Domain
func (ac AppContainer) TokenHelper() auth_domain.TokenHelper {
	if (ac.tokenHelper == nil) {
		tokenHelper := auth_domain.NewTokenHelper(
			jwt.SigningMethodHS256,
			[]byte(ac.configuration.JWT.SecretKey),
			ac.configuration.JWT.ExpiryDurationInDays,
		)
		ac.tokenHelper = &tokenHelper
	}
	return *ac.tokenHelper
}
func (ac AppContainer) RegisterUser(work uow.UnitOfWork) auth_domain.RegisterUser {
	return auth_domain.NewRegisterUser(
		work.UserRepo(),
		ac.TokenHelper(),
	)
}
func (ac AppContainer) LoginUser(work uow.UnitOfWork) auth_domain.LoginUser {
	return auth_domain.NewLoginUser(
		work.UserRepo(),
		ac.TokenHelper(),
	)
}
func (ac AppContainer) VerifyToken() auth_domain.VerifyToken {
	userRepo := user_repo.NewDefault(ac.dbConn)
	return auth_domain.NewVerifyToken(ac.TokenHelper(), userRepo)
}

func (ac AppContainer) CreateBook(work uow.UnitOfWork) book_domain.CreateBook {
	return book_domain.NewCreateBook(work.BookRepo(), work.ChapterRepo())
}
func (ac AppContainer) DeleteBook(work uow.UnitOfWork) book_domain.DeleteBook {
	return book_domain.NewDeleteBook(work.BookRepo())
}
func (ac AppContainer) UpdateBook(work uow.UnitOfWork) book_domain.UpdateBook {
	return book_domain.NewUpdateBook(work.BookRepo())
}

func (ac AppContainer) CreateChapterAndPath(work uow.UnitOfWork) chapter_domain.CreateChapterAndPath {
	return chapter_domain.NewCreateChapterAndPath(
		work.ChapterRepo(),
		work.ChapterPathRepo(),
		work.BookRepo(),
	)
}
func (ac AppContainer) UpdateChapter(work uow.UnitOfWork) chapter_domain.UpdateChapter {
	return chapter_domain.NewUpdateChapter(work.BookRepo(),work.ChapterRepo())
}
func (ac AppContainer) DeleteChapter(work uow.UnitOfWork) chapter_domain.DeleteChapter {
	return chapter_domain.NewDeleteChapter(work.ChapterRepo(),work.BookRepo())
}
func (ac AppContainer) CreatePathBetweenChapters(work uow.UnitOfWork) chapter_domain.CreatePathBetweenChapters {
	return chapter_domain.NewCreatePathBetweenChapters(
		work.ChapterRepo(),
		work.ChapterPathRepo(),
		work.BookRepo(),
	)
}
func (ac AppContainer) UpdatePathBetweenChapters(work uow.UnitOfWork) chapter_domain.UpdatePathBetweenChapters {
	return chapter_domain.NewUpdatePathBetweenChapters(
		work.ChapterRepo(),
		work.ChapterPathRepo(),
		work.BookRepo(),
	)
}

func (ac AppContainer) NavigateToChapter(work uow.UnitOfWork) game_domain.NavigateToChapter {
	return game_domain.NewNavigateToChapter(
		work.ChapterRepo(),
		work.ChapterPathRepo(),
	)
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
func (ac AppContainer) AuthMiddleware() auth_middleware.AuthMiddleware {
	userRepo := user_repo.NewDefault(ac.dbConn)
	getLoggedInUser := auth_domain.NewGetLoggedInUser(userRepo, ac.TokenHelper())

	return auth_middleware.New(getLoggedInUser)
}

/***********************************************************************************************/
/***********************************************************************************************/
