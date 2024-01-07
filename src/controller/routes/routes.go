package routes

import (
	"github.com/Emanueltyc/go-first-crud/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.RouterGroup) {
	router.GET("/getUserById/:userId", controller.FindUserByID)
	router.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	router.POST("/createUser", controller.CreateUser)
	router.PUT("/updateUser/:userId", controller.UpdateUser)
	router.DELETE("/deleteUser/:userId", controller.DeleteUser)
}
