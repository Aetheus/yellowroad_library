package routes

import (
	"github.com/gin-gonic/gin"
	"yellowroad_library/containers"
	"yellowroad_library/http/routes/book_crud_routes"
	"yellowroad_library/http/routes/chapter_crud_routes"
	"yellowroad_library/http/routes/story_routes"
	"yellowroad_library/http/routes/user_routes"
)

func registerRoutes(
	ginEngine *gin.Engine,
	container containers.Container,
){
	apiRouteGroup := ginEngine.Group("api")

	book_crud_routes.Register(apiRouteGroup,container)
	chapter_crud_routes.Register(apiRouteGroup,container)
	story_routes.Register(apiRouteGroup,container)
	user_routes.Register(apiRouteGroup,container)
}