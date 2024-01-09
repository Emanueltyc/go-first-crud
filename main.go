package main

import (
	"log"

	"github.com/Emanueltyc/go-first-crud/src/configuration/logger"
	"github.com/Emanueltyc/go-first-crud/src/controller"
	"github.com/Emanueltyc/go-first-crud/src/controller/routes"
	"github.com/Emanueltyc/go-first-crud/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user application")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
