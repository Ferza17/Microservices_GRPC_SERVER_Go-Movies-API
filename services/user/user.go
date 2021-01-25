package user

import (
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/domains/userDomain"
)

var (
	Services userServiceInterface = &userServiceStruct{}
)

type (
	userServiceStruct    struct{}
	userServiceInterface interface {
		Create(user userDomain.User) (*userDomain.User, error)
		GetUser(id int64) (*userDomain.User, error)
		Update(user userDomain.User) (*userDomain.User, error)
	}
)

func (u *userServiceStruct) Create(user userDomain.User) (*userDomain.User, error) {
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userServiceStruct) GetUser(id int64) (*userDomain.User, error) {
	result := &userDomain.User{
		Id: id,
	}

	if err := result.GetUser(); err != nil {
		return nil, err
	}

	return result, nil
}

func (u *userServiceStruct) Update(user userDomain.User) (*userDomain.User, error) {
	if err := user.Update(); err != nil {
		return nil, err
	}

	return &user, nil
}
