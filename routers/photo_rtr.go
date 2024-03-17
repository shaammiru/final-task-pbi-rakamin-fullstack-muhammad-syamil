package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupPhotoRouter(router *gin.Engine) {
	photoRouter := router.Group("/photos")
	{
		photoRouter.POST("/", createPhotoHandler)
		photoRouter.GET("/", getPhotosHandler)
		photoRouter.PUT("/:photoID", updatePhotoHandler)
		photoRouter.DELETE("/:photoID", deletePhotoHandler)
	}
}

func createPhotoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Create Photo Endpoint",
		"data":    nil,
	})
}

func getPhotosHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Photo Endpoint",
		"data":    nil,
	})

}

func updatePhotoHandler(c *gin.Context) {
	photoID := c.Param("photoID")
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Photo Endpoint",
		"data":    photoID,
	})
}

func deletePhotoHandler(c *gin.Context) {
	photoID := c.Param("photoID")
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Photo Endpoint",
		"data":    photoID,
	})
}
