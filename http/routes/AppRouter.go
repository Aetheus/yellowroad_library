package routes

import (
	"github.com/gin-gonic/gin"
	"yellowroad_library/containers"
	"net/http"
	"log"
	"time"
	"context"
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

func (this AppRouter) Shutdown (){
	srv := &http.Server{
		Addr:    ":8080",
		Handler: this.engine,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}