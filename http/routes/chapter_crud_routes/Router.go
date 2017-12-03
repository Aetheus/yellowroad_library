package chapter_crud_routes

import (
	"yellowroad_library/containers"

	"github.com/gin-gonic/gin"
)

func Register(
	router *gin.RouterGroup,
	container containers.Container,
) {
	chapterCrudHandlers := ChapterCrudHandlers{container:container}

	routesRequiringLogin := router.Group("", gin.HandlerFunc(container.GetAuthMiddleware()) )
	{
		routesRequiringLogin.POST("stories/:book_id/chapter", chapterCrudHandlers.CreateChapter)
		routesRequiringLogin.PUT("stories/:book_id/chapter/:chapter_id", chapterCrudHandlers.UpdateChapter)
		routesRequiringLogin.DELETE("stories/:book_id/chapter/:chapter_id", chapterCrudHandlers.DeleteChapter)
		routesRequiringLogin.POST("stories/:book_id/chapter/:chapter_id/paths", chapterCrudHandlers.CreatePathAwayFromThisChapter)
	}
}
