package routes

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	"yellowroad_library/containers"
	"yellowroad_library/http/routes/story_route"
	"yellowroad_library/http/routes/user_route"
)

func Init(container containers.Container) AppRouter {
	var ginEngine = gin.Default()
	ginEngine.Use(cors.New(cors.Config{
		//TODO : origins should be from a config!!
		AllowOrigins:     []string{"http://domainname.com", "http://localhost:3000"},

		AllowMethods:     []string{"PUT","PATCH","GET","POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	var r = newAppRouter(ginEngine, container)
	var portString = fmt.Sprintf(":%d", container.GetConfiguration().Web.Port)

	r.Route("/api/story", story_route.Register)
	r.Route("/api/users", user_route.Register)

	ginEngine.Run(portString)
	return r
}