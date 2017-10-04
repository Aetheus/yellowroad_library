package user_route

import (
	"yellowroad_library/containers"

	"github.com/gin-gonic/gin"
)

func Register(
	routerGroup *gin.RouterGroup,
	container containers.Container,
) {

	routerGroup.POST("/login", Login(container.GetAuthService()))
	routerGroup.POST("/register", SignUpHandler(container))

	// routerGroup.GET("/secure/secret", gin.HandlerFunc(authMiddleware), func(c *gin.Context) {

	// 	c.JSON(http.StatusOK, gin.H{"very": "secret"})
	// })

	// routerGroup.GET("/secure/not", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"not so": "secret hah"})
	// })

}
