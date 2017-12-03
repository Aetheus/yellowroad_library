package user_routes

import (
	"yellowroad_library/containers"

	"github.com/gin-gonic/gin"
)

func Register(
	router *gin.RouterGroup,
	container containers.Container,
) {
	userRouteHandlers := UserRouteHandlers{container:container}
	router.POST("users/login", userRouteHandlers.Login)
	router.POST("users/register", userRouteHandlers.SignUp)
	router.POST("users/verify", userRouteHandlers.VerifyToken)
}
