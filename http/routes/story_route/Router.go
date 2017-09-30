package story_route

import (
	"yellowroad_library/containers"

	"github.com/gin-gonic/gin"
)

func Register(
	routerGroup *gin.RouterGroup,
	container containers.Container,
) {

	//Book related
	{
		routerGroup.GET("/", FetchBooks(container.GetBookRepository()))
		routerGroup.GET("/:book_id", FetchSingleBook(container.GetBookRepository()))
		routesRequiringLogin := routerGroup.Group("", gin.HandlerFunc(container.GetAuthMiddleware()) )
		{
			routesRequiringLogin.POST("", CreateBook(container.GetAuthService(), container.GetBookService()))
			routesRequiringLogin.DELETE("/:book_id", DeleteBook(container.GetAuthService(), container.GetBookRepository(), container.GetBookService()))
		}
		routerGroup.PUT("/:book_id", UpdateBook(container.GetBookRepository()))
	}


	//Chapter/Story related
	{
		routerGroup.GET("/:book_id/:chapter_id", NavigateToSingleChapter(container.GetStoryService()))

	}

}
