package story_route

import (
	"yellowroad_library/containers"
	"yellowroad_library/http/routes/story_route/book_crud"
	"github.com/gin-gonic/gin"
)

func Register(
	routerGroup *gin.RouterGroup,
	container containers.Container,
) {

	//Book related
	{
		routerGroup.GET("/", book_crud.FetchBooks(container.GetBookRepository()))
		routerGroup.GET("/:book_id", book_crud.FetchSingleBook(container.GetBookRepository()))
		routesRequiringLogin := routerGroup.Group("", gin.HandlerFunc(container.GetAuthMiddleware()) )
		{
			routesRequiringLogin.POST("", book_crud.CreateBookHandler(container))
			routesRequiringLogin.DELETE("/:book_id", book_crud.DeleteBookHandler(container))
		}
		routerGroup.PUT("/:book_id", book_crud.UpdateBook(container.GetBookRepository()))
	}


	//Chapter/Story related
	{
		routerGroup.GET("/:book_id/:chapter_id", NavigateToSingleChapter(container.GetStoryService()))

	}

}
