package controller

import (
	"net/http"

	"github.com/Emanueltyc/go-first-crud/src/configuration/logger"
	"github.com/Emanueltyc/go-first-crud/src/configuration/validation"
	"github.com/Emanueltyc/go-first-crud/src/controller/model/request"
	"github.com/Emanueltyc/go-first-crud/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser Controller", zap.String("journey", "createUser"))

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))

		restError := validation.ValidateUserError(err)

		c.JSON(restError.Code, restError)
		return
	}

	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully", zap.String("journey", "createUser"))

	c.String(http.StatusOK, "")
}