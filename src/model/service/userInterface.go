package service

import (
	"github.com/Emanueltyc/go-first-crud/src/configuration/restERR"
	"github.com/Emanueltyc/go-first-crud/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *restERR.RestErr
	UpdateUser(string, model.UserDomainInterface) *restERR.RestErr
	FindUser(string) (*model.UserDomainInterface, *restERR.RestErr)
	DeleteUser(string) *restERR.RestErr
}
