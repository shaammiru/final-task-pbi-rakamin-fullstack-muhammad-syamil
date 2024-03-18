package app

import (
	"github.com/gin-gonic/gin"
	"github.com/shaammiru/task-5-pbi-fullstack-developer-muhammadsyamil/middlewares"
	"github.com/shaammiru/task-5-pbi-fullstack-developer-muhammadsyamil/routers"
)

func InitApp() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Use(middlewares.ServerErrorHandler())

	routers.SetupUserRouter(router)
	routers.SetupPhotoRouter(router)

	return router
}
