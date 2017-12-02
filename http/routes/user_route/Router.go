package user_route

import (
	"yellowroad_library/containers"

	"github.com/gin-gonic/gin"
)

func Register(
	routerGroup *gin.RouterGroup,
	container containers.Container,
) {
	workFactory := container.UnitOfWorkFactory()
	authServFactory := container.AuthServiceFactory()

	routerGroup.POST("/login", func (c *gin.Context){
		work := workFactory();
		authServ := authServFactory(work)
		Login(c,work, authServ)
	})
	routerGroup.POST("/register", func(c *gin.Context){
		work := workFactory();
		authServ := authServFactory(work)
		SignUp(c,work, authServ)
	})
	routerGroup.POST("/verify", func(c *gin.Context){
		work := workFactory();
		authServ := authServFactory(work)
		VerifyToken(c,work, authServ)
	})

	// routerGroup.GET("/secure/secret", gin.HandlerFunc(authMiddleware), func(c *gin.Context) {

	// 	c.JSON(http.StatusOK, gin.H{"very": "secret"})
	// })

	// routerGroup.GET("/secure/not", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"not so": "secret hah"})
	// })

}
