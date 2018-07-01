package routes

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	"yellowroad_library/containers"
	"yellowroad_library/utils/app_error"
)

func Init(container containers.AppContainer) app_error.AppError {
	var portString = fmt.Sprintf(":%d", container.Configuration().Web.Port)
	var ginEngine = gin.Default()

	ginEngine.Use(cors.New(cors.Config{
		AllowOrigins:     container.Configuration().Web.AllowOrigins,
		AllowMethods:     []string{"PUT","PATCH","GET","POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	fmt.Printf("CORS configured to allow the following origins: %s \n", container.Configuration().Web.AllowOrigins)

	ROUTES(ginEngine,container)

	ginEngine.Run(portString)

	return nil
}