package containers

import (
	"yellowroad_library/config"
	"yellowroad_library/database/repo/user_repo"
	"yellowroad_library/http/middleware/auth_middleware"
	"yellowroad_library/services/auth_serv"
	"yellowroad_library/services/token_serv"
	"yellowroad_library/services/book_serv"
	"yellowroad_library/database/repo/book_repo"
	"yellowroad_library/services/story_serv"
	"yellowroad_library/database/repo/chapter_repo"
	"yellowroad_library/database/repo/chapterpath_repo"
	"yellowroad_library/database/repo/uow"
	"yellowroad_library/services/auth_serv/user_registration_serv"
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
		UserRegistrationService(work *uow.UnitOfWork,autocommit bool) user_registration_serv.UserRegistrationService

	GetTokenService() token_serv.TokenService
	GetBookService() book_serv.BookService
	GetStoryService() story_serv.StoryService

	//repo
	GetUserRepository() user_repo.UserRepository
	GetBookRepository() book_repo.BookRepository
	GetChapterRepository() chapter_repo.ChapterRepository
	GetChapterPathRepository() chapterpath_repo.ChapterPathRepository
	UnitOfWork() uow.UnitOfWork

	//middleware
	GetAuthMiddleware() auth_middleware.AuthMiddleware

	//configuration
	GetConfiguration() config.Configuration
}
