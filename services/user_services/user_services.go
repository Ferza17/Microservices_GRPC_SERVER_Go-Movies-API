package user_services

import (
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/domains/user_domain"
)

var (
	Services userServiceInterface = &userServiceStruct{}
)

type (
	userServiceStruct    struct{}
	userServiceInterface interface {
		Create(user user_domain.User) (*user_domain.User, error)
		GetUser(id int64) (*user_domain.User, error)
		Update(user user_domain.User) (*user_domain.User, error)
	}
)

func (u *userServiceStruct) Create(user user_domain.User) (*user_domain.User, error) {
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userServiceStruct) GetUser(id int64) (*user_domain.User, error) {
	result := &user_domain.User{
		Id: id,
	}

	if err := result.GetUser(); err != nil {
		return nil, err
	}

	return result, nil
}

func (u *userServiceStruct) Update(user user_domain.User) (*user_domain.User, error) {
	if err := user.Update(); err != nil {
		return nil, err
	}

	return &user, nil
}
