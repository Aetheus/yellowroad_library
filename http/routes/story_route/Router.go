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
	chapterServFactory := container.ChapterServiceFactory()

	//Book CRUD related
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

	//Chapter CRUD related
	{
		routesRequiringLogin := routerGroup.Group("", gin.HandlerFunc(container.GetAuthMiddleware()) )
		{
			routesRequiringLogin.POST("/:book_id/chapter", func (c *gin.Context){
				work := workFactory()
				CreateChapter(c,work, authServFactory(work), chapterServFactory(work))
			})

			routesRequiringLogin.PUT("/:book_id/chapter/:chapter_id", func (c *gin.Context){
				work := workFactory()
				UpdateChapter(c,work, authServFactory(work), chapterServFactory(work))
			})

			routesRequiringLogin.DELETE("/:book_id/chapter/:chapter_id", func (c *gin.Context){
				work := workFactory()
				DeleteChapter(c,work, authServFactory(work), chapterServFactory(work))
			})

			routesRequiringLogin.POST("/:book_id/chapter/:chapter_id/paths", func(c *gin.Context){
				work := workFactory()
				CreatePathAwayFromThisChapter(c, work, authServFactory(work), chapterServFactory(work))
			})
		}
	}




	//Game related
	{
		routerGroup.GET("/:book_id/chapter/:chapter_id/game", func(c *gin.Context){
			work := workFactory()
			NavigateToSingleChapter(c,work,storyServFactory(work))
		})
	}

}
