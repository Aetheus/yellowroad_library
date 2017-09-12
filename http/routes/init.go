package routes

import (
	"fmt"
	"yellowroad_library/containers"
	"yellowroad_library/http/routes/bookRoute"

	"github.com/gin-gonic/gin"
)

func Init(container containers.Container) {
	var ginEngine = gin.Default()
	var r = newAppRouter(ginEngine, container)
	var portString = fmt.Sprintf(":%d", container.GetConfiguration().Web.Port)

	r.Route("/api/books", bookRoute.Register)

	ginEngine.Run(portString)
}