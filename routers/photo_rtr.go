package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/shaammiru/task-5-pbi-fullstack-developer-muhammadsyamil/controllers"
	"github.com/shaammiru/task-5-pbi-fullstack-developer-muhammadsyamil/helpers"
	"github.com/shaammiru/task-5-pbi-fullstack-developer-muhammadsyamil/models"
	"net/http"
	"strconv"
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
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"data":    nil,
			"error":   "Token is required",
		})
		return
	}

	claims, err := helpers.GetTokenClaims(token[7:])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	var photoData models.PhotoCreate
	if err := c.ShouldBindJSON(&photoData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Create photo failed",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	err = helpers.ValidateStruct(photoData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Create photo failed",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	user, err := controllers.GetUserByID(strconv.Itoa(int(claims.ID)))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"data":    nil,
			"error":   "User not found",
		})
		return
	}

	newPhoto, err := controllers.CreatePhoto(models.Photo{
		Title:    photoData.Title,
		Caption:  photoData.Caption,
		PhotoURL: photoData.PhotoURL,
		UserID:   user.ID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Create photo failed",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Create Photo Endpoint",
		"data":    newPhoto,
	})
}

func getPhotosHandler(c *gin.Context) {
	var photos []models.Photo

	photos, err := controllers.ListPhotos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Get photos failed",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get Photo Endpoint",
		"data":    photos,
	})

}

func updatePhotoHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"data":    nil,
			"error":   "Token is required",
		})
		return
	}

	claims, err := helpers.GetTokenClaims(token[7:])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	photoID := c.Param("photoID")

	var photoData models.PhotoUpdate
	if err := c.ShouldBindJSON(&photoData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Update photo failed",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	err = helpers.ValidateStruct(photoData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Update photo failed",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	photo, err := controllers.GetPhotoByID(photoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Update photo failed",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	if photo.UserID != claims.ID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"data":    nil,
			"error":   "User not permitted to update other user's photo",
		})
		return
	}

	updatedPhoto, err := controllers.UpdatePhotoByID(photoID, models.Photo{
		Title:    photoData.Title,
		Caption:  photoData.Caption,
		PhotoURL: photoData.PhotoURL,
		UserID:   photo.UserID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Update photo failed",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Update photo success",
		"data":    updatedPhoto,
	})
}

func deletePhotoHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"data":    nil,
			"error":   "Token is required",
		})
		return
	}

	claims, err := helpers.GetTokenClaims(token[7:])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	photoID := c.Param("photoID")
	photo, err := controllers.GetPhotoByID(photoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Update photo failed",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	if photo.UserID != claims.ID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"data":    nil,
			"error":   "User not permitted to delete other user's photo",
		})
		return
	}

	err = controllers.DeletePhotoByID(photoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Delete Photo failed",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Photo Endpoint",
		"data":    photoID,
	})
}
