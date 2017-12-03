package routes

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	"yellowroad_library/containers"
)

func Init(container containers.Container) error {
	var portString = fmt.Sprintf(":%d", container.GetConfiguration().Web.Port)
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

	registerRoutes(ginEngine,container)
	ginEngine.Run(portString)

	return nil
}