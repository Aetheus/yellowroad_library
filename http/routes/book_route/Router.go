package book_route

import (
	"yellowroad_library/containers"

	"github.com/gin-gonic/gin"
)

func Register(
	routerGroup *gin.RouterGroup,
	container containers.Container,
) {

	routerGroup.GET("/", FetchBooks(container.GetBookRepository()))
	routerGroup.GET("/:book_id", FetchSingleBook(container.GetBookRepository()))

	routesRequiringLogin := routerGroup.Group("", gin.HandlerFunc(container.GetAuthMiddleware()) )
	{
		routesRequiringLogin.POST("", CreateBook(container.GetAuthService(), container.GetBookService()))
		routesRequiringLogin.DELETE("/:book_id", DeleteBook(container.GetAuthService(), container.GetBookRepository(), container.GetBookService()))
	}

	routerGroup.PUT("/:book_id", UpdateBook(container.GetBookRepository()))

	// routerGroup.GET("/secure/secret", gin.HandlerFunc(authMiddleware), func(c *gin.Context) {

	// 	c.JSON(http.StatusOK, gin.H{"very": "secret"})
	// })

	// routerGroup.GET("/secure/not", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"not so": "secret hah"})
	// })

}
