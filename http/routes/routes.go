package routes

import (
	"github.com/gin-gonic/gin"
	"yellowroad_library/containers"
	"yellowroad_library/http/routes/book_crud_routes"
	"yellowroad_library/http/routes/chapter_crud_routes"
	"yellowroad_library/http/routes/story_routes"
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
		public_api.GET("stories/:book_id", bookCrudHandlers.FetchSingleBook)
		auth_api.POST("stories", bookCrudHandlers.CreateBook)
		auth_api.PUT("stories/:book_id", bookCrudHandlers.UpdateBook)
		auth_api.DELETE("stories/:book_id", bookCrudHandlers.DeleteBook)
	}

	chapterCrudHandlers := chapter_crud_routes.ChapterCrudHandlers{container}
	{
		auth_api.POST("stories/:book_id/chapters", chapterCrudHandlers.CreateChapter)
		public_api.GET("stories/:book_id/chapters/:chapter_id", chapterCrudHandlers.FetchSingleChapter)
		auth_api.PUT("stories/:book_id/chapters/:chapter_id", chapterCrudHandlers.UpdateChapter)
		auth_api.DELETE("stories/:book_id/chapters/:chapter_id", chapterCrudHandlers.DeleteChapter)
		auth_api.POST("stories/:book_id/chapters/:chapter_id/paths", chapterCrudHandlers.CreatePathAwayFromThisChapter)
	}

	storyHandlers := story_routes.StoryHandlers{container}
	{
		public_api.POST("stories/:book_id/chapters/:chapter_id/game", storyHandlers.NavigateToSingleChapter)
	}

	userRouteHandlers := user_routes.UserRouteHandlers{container}
	{
		public_api.POST("users/login", userRouteHandlers.Login)
		public_api.POST("users/register", userRouteHandlers.SignUp)
		public_api.POST("users/verify", userRouteHandlers.VerifyToken)
	}
}
