package routes

import (
	"github.com/gin-gonic/gin"
	"yellowroad_library/containers"
)

type AppRouter struct {
	engine *gin.Engine
	container containers.Container
}

func newAppRouter(engine *gin.Engine, container containers.Container) AppRouter {
	return AppRouter{
		engine, container,
	}
}

func (this AppRouter) Route (
	routePath string,
	registrationFunction func (
		*gin.RouterGroup,
		containers.Container,
	),
) {
	registrationFunction(this.engine.Group(routePath), this.container)
}
