package service

import (
	"github.com/devsouzx/crud-go/src/configuration/rest_err"
	"github.com/devsouzx/crud-go/src/model"
	"github.com/devsouzx/crud-go/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserServices(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUserServices(string, model.UserDomainInterface) *rest_err.RestErr
	FindUserByIDServices(string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailServices(string) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUserServices(string) *rest_err.RestErr
}
