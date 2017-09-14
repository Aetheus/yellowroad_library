package routes

import (
	"fmt"
	"yellowroad_library/containers"
	"yellowroad_library/http/routes/bookRoute"

	"github.com/gin-gonic/gin"
	"yellowroad_library/http/routes/userRoute"
)

func Init(container containers.Container) {
	var ginEngine = gin.Default()
	var r = newAppRouter(ginEngine, container)
	var portString = fmt.Sprintf(":%d", container.GetConfiguration().Web.Port)

	r.Route("/api/books", bookRoute.Register)
	r.Route("/api/users", userRoute.Register)

	ginEngine.Run(portString)
}