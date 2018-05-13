package routes

import (
	"github.com/gin-gonic/gin"
	"yellowroad_library/containers"
	"yellowroad_library/http/routes/book_crud_routes"
	"yellowroad_library/http/routes/chapter_crud_routes"
	"yellowroad_library/http/routes/game_routes"
	"yellowroad_library/http/routes/user_routes"
)


func ROUTES(
	ginEngine *gin.Engine,
	container containers.Container,
){
	// Group to prefix all routes with "api"
	public_api := ginEngine.Group("api")

	// Group that will be used by all routes that need a Auth Middleware
	// (i.e: routes that require the user to be logged in)
	auth_api := public_api.Group("", gin.HandlerFunc(container.GetAuthMiddleware()))

	bookCrudHandlers := book_crud_routes.BookCrudHandlers{container}
	{
		public_api.GET("stories", bookCrudHandlers.FetchBooks)
		auth_api.POST("stories", bookCrudHandlers.CreateBook)

		public_api.GET("stories/:book_id", bookCrudHandlers.FetchSingleBook)
		auth_api.PUT("stories/:book_id", bookCrudHandlers.UpdateBook)
		auth_api.DELETE("stories/:book_id", bookCrudHandlers.DeleteBook)
	}

	gameHandlers := game_routes.GameHandlers{container}
	{
		public_api.POST("stories/:book_id/chapters/:chapter_id/game", gameHandlers.NavigateToSingleChapter)
	}

	chapterCrudHandlers := chapter_crud_routes.ChapterCrudHandlers{container}
	{
		auth_api.POST("stories/:book_id/chapters", chapterCrudHandlers.CreateChapter)

		public_api.GET("stories/:book_id/chapters", chapterCrudHandlers.FetchChaptersIndex)
		public_api.GET("stories/:book_id/chapters/:chapter_id", chapterCrudHandlers.FetchSingleChapter)
		auth_api.PUT("stories/:book_id/chapters/:chapter_id", chapterCrudHandlers.UpdateChapter)
		auth_api.DELETE("stories/:book_id/chapters/:chapter_id", chapterCrudHandlers.DeleteChapter)

		auth_api.POST("stories/:book_id/chapters/:chapter_id/paths", chapterCrudHandlers.CreatePathAwayFromThisChapter)
	}

	userRouteHandlers := user_routes.UserRouteHandlers{container}
	{
		public_api.POST("users", userRouteHandlers.SignUp)

		public_api.POST("tokens", userRouteHandlers.Login)
		public_api.POST("tokens/verify", userRouteHandlers.VerifyToken)
	}
}
