package service

import (
	"github.com/Emanueltyc/go-first-crud/src/configuration/logger"
	"github.com/Emanueltyc/go-first-crud/src/configuration/restERR"
	"github.com/Emanueltyc/go-first-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *restERR.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	return nil
}
