package app

import (
	"github.com/gin-gonic/gin"
	"github.com/shaammiru/task-5-pbi-fullstack-developer-muhammadsyamil/middlewares"
	"github.com/shaammiru/task-5-pbi-fullstack-developer-muhammadsyamil/routers"
)

func InitApp() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	app := gin.Default()

	app.Use(middlewares.ServerPanicHandler())

	routers.SetupUserRouter(app)
	routers.SetupPhotoRouter(app)

	return app
}
