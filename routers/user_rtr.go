package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/shaammiru/final-task-pbi-rakamin-fullstack-muhammad-syamil/controllers"
	"github.com/shaammiru/final-task-pbi-rakamin-fullstack-muhammad-syamil/helpers"
	"github.com/shaammiru/final-task-pbi-rakamin-fullstack-muhammad-syamil/models"
	"net/http"
)

func SetupUserRouter(router *gin.Engine) {
	userRouter := router.Group("/users")
	{
		userRouter.GET("/", listUserHandler)
		userRouter.POST("/register", registerHandler)
		userRouter.POST("/login", loginHandler)
		userRouter.PUT("/:userID", updateUserHandler)
		userRouter.DELETE("/:userID", deleteUserHandler)
	}
}

func registerHandler(c *gin.Context) {
	var userData models.UserRegister
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Register failed",
			"data":    nil,
			"error":   "Invalid JSON format, check your request body",
		})
		return
	}

	err := helpers.ValidateStruct(userData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Register failed",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	userData.Password, err = helpers.HashPassword(userData.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Register failed",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	newUser, err := controllers.CreateUser(models.User{
		Username: userData.Username,
		Email:    userData.Email,
		Password: userData.Password,
	})
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
	var userData models.UserLogin
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Login failed",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	err := helpers.ValidateStruct(userData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Login failed",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	existUser, err := controllers.GetUserByEmail(userData.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Login failed",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	err = helpers.ComparePassword(existUser.Password, userData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Login failed",
			"data":    nil,
			"error":   "Invalid password",
		})
		return
	}

	token, err := helpers.GenerateToken(models.User{
		ID:       existUser.ID,
		Username: existUser.Username,
		Email:    existUser.Email,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Login failed",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login success",
		"token":   token,
	})
}

func listUserHandler(c *gin.Context) {
	users, err := controllers.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "List users failed",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "List users success",
		"data":    users,
	})
}

func updateUserHandler(c *gin.Context) {
	userID := c.Param("userID")

	var userData models.UserUpdate
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Update user failed",
			"data":    nil,
			"error":   "Invalid JSON format, check your request body",
		})
		return
	}

	err := helpers.ValidateStruct(userData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Update user failed",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	updatedUser, err := controllers.UpdateUserByID(userID, models.User{
		Username: userData.Username,
		Email:    userData.Email,
		Password: userData.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Update user failed",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Update user success",
		"data":    updatedUser,
	})
}

func deleteUserHandler(c *gin.Context) {
	userID := c.Param("userID")

	err := controllers.DeleteUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Delete User failed",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete User success",
	})
}
