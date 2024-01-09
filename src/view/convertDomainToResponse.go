package view

import (
	"github.com/Emanueltyc/go-first-crud/src/controller/model/response"
	"github.com/Emanueltyc/go-first-crud/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
