package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/shaammiru/task-5-pbi-fullstack-developer-muhammadsyamil/models"
	"net/http"

	"github.com/shaammiru/task-5-pbi-fullstack-developer-muhammadsyamil/controllers"
)

func SetupUserRouter(router *gin.Engine) {
	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", registerHandler)
		userRouter.POST("/login", loginHandler)
		userRouter.PUT("/:userID", updateUserHandler)
		userRouter.DELETE("/:userID", deleteUserHandler)
	}
}

func registerHandler(c *gin.Context) {
	var userData models.User
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Register failed",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	newUser, err := controllers.CreateUser(userData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Register failed",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Register success",
		"data":    newUser,
	})
}

func loginHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Login Endpoint",
		"data":    nil,
	})
}

func updateUserHandler(c *gin.Context) {
	userID := c.Param("userID")
	c.JSON(http.StatusOK, gin.H{
		"message": "Update User Endpoint",
		"data":    userID,
	})
}

func deleteUserHandler(c *gin.Context) {
	userID := c.Param("userID")
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete User Endpoint",
		"data":    userID,
	})
}
