package containers

import (
	"yellowroad_library/config"
	"yellowroad_library/http/middleware/auth_middleware"
	"yellowroad_library/services/auth_serv"
	"yellowroad_library/services/token_serv"
	"yellowroad_library/services/book_serv"
	"yellowroad_library/services/game_serv"
	"yellowroad_library/database/repo/uow"
	"yellowroad_library/services/chapter_serv"
)

type Container interface {
	//services
	TokenService() token_serv.TokenService
	AuthService(work uow.UnitOfWork) auth_serv.AuthService
	BookService(work uow.UnitOfWork) book_serv.BookService
	ChapterService(work uow.UnitOfWork) chapter_serv.ChapterService
	StoryService(work uow.UnitOfWork) game_serv.GameService

	//Unit of Work
	UnitOfWork() uow.UnitOfWork

	//middleware
	GetAuthMiddleware() auth_middleware.AuthMiddleware

	//configuration
	GetConfiguration() config.Configuration
}
