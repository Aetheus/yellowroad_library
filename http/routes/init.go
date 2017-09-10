package routes

import (
	"fmt"
	"yellowroad_library/containers"
	"yellowroad_library/http/routes/BookRoute"

	"github.com/gin-gonic/gin"
)

func Init(container containers.Container) {
	var r = gin.Default()

	BookRoute.Register(r.Group("/api/books"), container)

	var portString = fmt.Sprintf(":%d", container.GetConfiguration().Web.Port)
	r.Run(portString)
}
