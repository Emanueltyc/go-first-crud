package model

import (
	"github.com/Emanueltyc/go-first-crud/src/configuration/logger"
	"github.com/Emanueltyc/go-first-crud/src/configuration/restERR"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *restERR.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	return nil
}
