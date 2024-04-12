package app

import (
	"github.com/gin-gonic/gin"
	"github.com/shaammiru/final-task-pbi-rakamin-fullstack-muhammad-syamil/middlewares"
	"github.com/shaammiru/final-task-pbi-rakamin-fullstack-muhammad-syamil/routers"
)

func InitApp() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	app := gin.Default()

	app.Use(middlewares.ServerPanicHandler())

	routers.SetupUserRouter(app)
	routers.SetupPhotoRouter(app)

	return app
}
