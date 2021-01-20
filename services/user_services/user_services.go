package user_services

import (
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/domains/user_domain"
)

var (
	UserServices userServiceInterface = &userServiceStruct{}
)

type (
	userServiceStruct    struct{}
	userServiceInterface interface {
		Create(user user_domain.User) (*user_domain.User, error)
	}
)

func (u *userServiceStruct) Create(user user_domain.User) (*user_domain.User, error) {
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
