package book_crud_routes

import (
	"yellowroad_library/containers"

	"github.com/gin-gonic/gin"
)

func Register(
	router *gin.RouterGroup,
	container containers.Container,
) {
	bookHandlers := BookCrudHandlers{container:container}

	router.GET("stories/", bookHandlers.FetchBooks)
	router.GET("stories/:book_id", bookHandlers.FetchSingleBook)

	authRoutes := router.Group("", gin.HandlerFunc(container.GetAuthMiddleware()))
	{
		authRoutes.POST("stories", bookHandlers.CreateBook)
		authRoutes.PUT("stories/:book_id", bookHandlers.UpdateBook)
		authRoutes.DELETE("stories/:book_id", bookHandlers.DeleteBook)
	}
}
