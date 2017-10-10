package story_route

import (
	"yellowroad_library/containers"

	"github.com/gin-gonic/gin"
)

func Register(
	routerGroup *gin.RouterGroup,
	container containers.Container,
) {
	workFactory := container.UnitOfWorkFactory()
	bookServFactory := container.BookServiceFactory()
	authServFactory := container.AuthServiceFactory()
	storyServFactory := container.StoryServiceFactory()

	//Book related
	{
		routerGroup.GET("/", func (c *gin.Context){
			FetchBooks(c, workFactory())
		})

		routerGroup.GET("/:book_id", func(c *gin.Context) {
			FetchSingleBook(c,workFactory())
		})

		routesRequiringLogin := routerGroup.Group("", gin.HandlerFunc(container.GetAuthMiddleware()) )
		{
			//func(c *gin.Context) { }
			routesRequiringLogin.POST("",func(c *gin.Context) {
				work := workFactory()
				CreateBook(c, work, authServFactory(work), bookServFactory(work) )
			})

			routesRequiringLogin.DELETE("/:book_id", func(c *gin.Context){
				work := workFactory()
				DeleteBook(c, work, authServFactory(work), bookServFactory(work) )
			})

			routerGroup.PUT("/:book_id", func(c *gin.Context){
				work := workFactory()
				UpdateBook(c, work, authServFactory(work), bookServFactory(work) )
			})
		}

	}


	//Chapter/Story related
	{
		routerGroup.GET("/:book_id/:chapter_id", func(c *gin.Context){
			work := workFactory()
			NavigateToSingleChapter(c,work,storyServFactory(work))
		})
	}

}
