package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/Emanueltyc/go-first-crud/src/configuration/restERR"
)

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &UserDomain{
		email,
		password,
		name,
		age,
	}
}

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *restERR.RestErr
	UpdateUser(string) *restERR.RestErr
	FindUser(string) (*UserDomain, *restERR.RestErr)
	DeleteUser(string) *restERR.RestErr
}
