package routes

import (
	"github.com/Emanueltyc/go-first-crud/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.RouterGroup, userController controller.UserControllerInterface) {
	router.GET("/getUserById/:userId", userController.FindUserByID)
	router.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	router.POST("/createUser", userController.CreateUser)
	router.PUT("/updateUser/:userId", userController.UpdateUser)
	router.DELETE("/deleteUser/:userId", userController.DeleteUser)
}
